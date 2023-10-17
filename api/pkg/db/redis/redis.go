package redis

import (
	"log"
	"time"

	"context"
	"github.com/go-redis/redis/v8"
)

var _ RedisFactory = (*RedisServ)(nil)

type RedisFactory interface {
	GetRedisClient() *redis.Client
	Close() error
	GetValue(key string) (string, error)
	SetNx(key string, value interface{}, t time.Duration) error
	Del(key string) error
	Keys(pattern string) []string
	Exist(key string) bool
}

type RedisServ struct {
	_redisClient *redis.Client
}

func InitRds(redisip, pw string, db int) RedisFactory {
	client := redis.NewClient(&redis.Options{
		Addr:         redisip,
		Password:     pw,
		DB:           db,
		IdleTimeout:  20 * time.Second,
		PoolSize:     20,
		MinIdleConns: 10,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("[INIT] init redis database failed %v", err)
		return nil
	}
	return &RedisServ{
		_redisClient: client,
	}
}

func (r *RedisServ) GetRedisClient() *redis.Client {
	return r._redisClient
}

func (r *RedisServ) Close() error {
	return r._redisClient.Close()
}

func (r *RedisServ) GetValue(key string) (string, error) {
	return r._redisClient.Get(context.Background(), key).Result()
}

func (r *RedisServ) SetNx(key string, value interface{}, t time.Duration) error {
	_, err := r._redisClient.Set(context.Background(), key, value, t).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisServ) Del(key string) error {
	return r._redisClient.Del(context.Background(), key).Err()
}

// 获取所有key
func (r *RedisServ) Keys(pattern string) []string {
	values, _ := r._redisClient.Keys(context.Background(), pattern+"*").Result()
	return values
}

// 判断key是否存在
func (r *RedisServ) Exist(key string) bool {
	return r.Exist(key)
}
