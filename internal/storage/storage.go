package storage

import (
	"context"
	"jordanmckoy/guardian/db"
	"log"
	"os"

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

	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))

	if err != nil {
		log.Fatalf("Unable to connect to redis: %v", err)
	}

	RedisClient = redis.NewClient(opts)

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		log.Fatalf("Unable to connect to redis: %v", err)
		return
	}
}
