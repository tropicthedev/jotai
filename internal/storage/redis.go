package storage

import (
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client = rdb
}

func GetClient() *redis.Client {
	return client
}
