package geniuslurker

import (
	"sync"

	"github.com/go-redis/redis"
)

type singleton struct {
}

const redisURL = "localhost:6379"
const redisPassowrd = "" // no password set
const redisDB = 0        // use default DB

var redisClient *redis.Client

var once sync.Once

// GetRedisClient returns instance of a Redis client
func GetRedisClient() *redis.Client {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: redisPassowrd,
			DB:       redisDB,
		})
	})
	return redisClient
}
