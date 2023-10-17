package lock

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"mm-ndj/config"
	"mm-ndj/pkg/utils/rand"
	"strconv"
	"sync/atomic"
)

var (
	// 上锁lua
	lockCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2]);
    return "OK";
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2]);
end`

	// 释放锁lua
	delCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1]);
else
    return 0;
end`

	// 默认超时时间，防止死锁
	tolerance       = 500 // milliseconds
	millisPerSecond = 1000
)

type RedisLock struct {
	cli *redis.Client
	// 超时时间
	seconds uint32
	// 锁value，防止锁被别人获取到
	id string
}

func NewRedisLock(cli *redis.Client, seconds uint32) RedisLock {
	return RedisLock{
		cli:     cli,
		id:      rand.RandString(16), // 随机字符串，作为锁持有者的标识，防止锁被非持有者释放
		seconds: seconds,
	}
}

// Lock 上锁
func (lock *RedisLock) Lock(key string) (bool, error) {
	// 获取过期时间
	seconds := atomic.LoadUint32(&lock.seconds)
	// 默认锁过期时间为500ms，防止死锁
	resp, err := lock.cli.Eval(lock.cli.Context(),
		lockCommand, []string{key}, lock.id,
		strconv.Itoa(int(seconds)*millisPerSecond+tolerance)).Result()

	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		config.Logger.Error("获得锁时发生异常：", zap.Error(err))
		return false, err
	} else if resp == nil {
		return false, nil
	}

	reply, ok := resp.(string)
	if ok && reply == "OK" {
		return true, nil
	}

	config.Logger.Error("获得锁时收到redis的未知回复：", zap.Error(err))
	return false, nil
}

// UnLock 释放锁
func (lock *RedisLock) UnLock(key string) (bool, error) {
	resp, err := lock.cli.Eval(lock.cli.Context(), delCommand, []string{key}, lock.id).Result()
	if err != nil {
		config.Logger.Error("释放锁时发生异常：", zap.Error(err))
		return false, err
	}

	reply, ok := resp.(int64)
	if !ok {
		return false, nil
	}

	return reply == 1, nil
}

// SetExpire 自定义锁过期时间
// 需要注意的是需要在Lock()之前调用
// 不然默认为500ms自动释放
func (lock *RedisLock) SetExpire(seconds int) {
	atomic.StoreUint32(&lock.seconds, uint32(seconds))
}
