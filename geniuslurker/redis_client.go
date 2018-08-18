package geniuslurker

import (
	"encoding/json"

	"github.com/AlexanderYAPPO/geniuslurker/datastructers"
	"github.com/go-redis/redis"
)

const redisURL = "localhost:6379"
const redisPassowrd = "" // no password set
const redisDB = 0        // use default DB

//RedisClient is a real Redis client
type RedisClient struct {
	redisClient *redis.Client
}

//Exists checks if exist
func (client *RedisClient) Exists(key string) bool {
	result, err := client.redisClient.Exists(key).Result()
	if err != nil {
		ErrorLogger.Panicln("Error accessing redis", err)
	}
	return result != 0
}

//SearchResultsRPushJSON pushes object of SearchResult inte Redis array value
func (client *RedisClient) SearchResultsRPushJSON(key string, value datastructers.SearchResult) {
	valueB, _ := json.Marshal(value)
	_, err := client.redisClient.RPush(key, valueB).Result()
	if err != nil {
		ErrorLogger.Panicln("Error accessing redis", err)
	}
}

//Del deletes value
func (client *RedisClient) Del(key string) {
	_, err := client.redisClient.Del(key).Result()
	if err != nil {
		ErrorLogger.Panicln("Error accessing redis", err)
	}
}

//LLen get the size of an array balue
func (client *RedisClient) LLen(key string) int64 {
	size, err := client.redisClient.LLen(key).Result()
	if err != nil {
		ErrorLogger.Panicln("Error accessing redis", err)
	}
	return size
}

//SearchResultsIndexJSON gets ith value of an array
func (client *RedisClient) SearchResultsIndexJSON(key string, index int64) datastructers.SearchResult {
	searchResultB, err := client.redisClient.LIndex(key, index).Bytes()
	if err != nil {
		ErrorLogger.Panicln("Error accessing redis", err)
	}
	var searchResult datastructers.SearchResult
	json.Unmarshal(searchResultB, &searchResult)
	return searchResult
}

//NewRedisClient returns RedisClient object
func NewRedisClient() *RedisClient {
	newClient := &RedisClient{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: redisPassowrd,
			DB:       redisDB,
		}),
	}
	return newClient
}
