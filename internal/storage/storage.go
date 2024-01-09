package storage

import (
	"context"
	"jordanmckoy/guardian/db"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient  *redis.Client
	PrismaClient *db.PrismaClient
)

func init() {
	ctx := context.Background()

	PrismaClient = db.NewClient()

	if err := PrismaClient.Prisma.Connect(); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		log.Fatalf("Unable to connect to redis: %v", err)
		return
	}
}
