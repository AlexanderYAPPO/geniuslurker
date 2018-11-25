package mocks

import (
	"github.com/AlexanderYAPPO/geniuslurker/datastructers"
)

const redisURL = "localhost:6379"
const redisPassowrd = "" // no password set
const redisDB = 0        // use default DB

//RedisClientMock is a real Redis client
type RedisClientMock struct {
	storage map[string][]datastructers.SearchResult
}

//Exists checks if exist
func (client *RedisClientMock) Exists(key string) bool {
	_, ok := client.storage[key]
	return ok
}

//SearchResultsRPushJSON pushes object of SearchResult inte Redis array value
func (client *RedisClientMock) SearchResultsRPushJSON(key string, value datastructers.SearchResult) {
	v, ok := client.storage[key]
	if !ok {
		client.storage[key] = []datastructers.SearchResult{value}
		return
	}
	v = append(v, value)
	client.storage[key] = v
}

//Del deletes value
func (client *RedisClientMock) Del(key string) {
	delete(client.storage, key)
}

//LLen get the size of an array balue
func (client *RedisClientMock) LLen(key string) int64 {
	v, _ := client.storage[key]
	return int64(len(v))
}

//SearchResultsIndexJSON gets ith value of an array
func (client *RedisClientMock) SearchResultsIndexJSON(key string, index int64) datastructers.SearchResult {
	v, _ := client.storage[key]
	return v[index]
}

//NewRedisClient returns RedisClient object
func NewRedisClient() *RedisClientMock {
	newClient := &RedisClientMock{make(map[string][]datastructers.SearchResult)}
	return newClient
}
