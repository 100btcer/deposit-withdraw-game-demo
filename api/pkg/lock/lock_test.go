package lock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRedisLock(t *testing.T) {
	redisLock := NewRedisLock(redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 10,
		Password:     "",
	}), 5)
	assert.NotEmpty(t, redisLock.cli)
	result, err := redisLock.cli.Ping(context.Background()).Result()
	assert.Empty(t, err)
	assert.Equal(t, result, "PONG")
}

func TestRedisLock_Lock(t *testing.T) {
	cli := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 10,
		Password:     "",
	})

	lockKey := "lock:key:xxxx"
	// lock1
	firstLock := NewRedisLock(cli, 5)
	firstAcquire, err := firstLock.Lock(lockKey)
	assert.Nil(t, err)
	assert.True(t, firstAcquire)

	// 重入
	firstAcquire, err = firstLock.Lock(lockKey)
	assert.Nil(t, err)
	assert.True(t, firstAcquire)

	// lock2
	secondLock := NewRedisLock(cli, 5)
	againAcquire, err := secondLock.Lock(lockKey)
	assert.Nil(t, err)
	assert.False(t, againAcquire)

	release, err := firstLock.UnLock(lockKey)
	assert.Nil(t, err)
	assert.True(t, release)

	endAcquire, err := secondLock.Lock(lockKey)
	assert.Nil(t, err)
	assert.True(t, endAcquire)
}
