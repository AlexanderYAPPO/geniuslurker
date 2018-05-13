package geniuslurker

import (
	"sync"
)

var redisClient RedisClientI

var onceRedisClient sync.Once

// RedisClientI represents interface of a client for accessing Redis
type RedisClientI interface {
	Exists(key string) bool
	SearchResultsRPushJSON(key string, value SearchResult)
	Del(key string)
	LLen(key string) int64
	SearchResultsIndexJSON(key string, index int64) SearchResult
}

// GetRedisClient returns instance of a Redis client
func GetRedisClient() RedisClientI {
	onceRedisClient.Do(func() {
		redisClient = NewRedisClient()
	})
	return redisClient
}
