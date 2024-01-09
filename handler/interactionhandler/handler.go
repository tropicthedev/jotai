package interactionhandler

import (
	"jordanmckoy/guardian/config"
	"jordanmckoy/guardian/db"
	"jordanmckoy/guardian/internal/storage"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	redis      *redis.Client
	prisma     *db.PrismaClient
	config     *config.Discord
	handlers   []func() // we need handlers func in order to deregister them later
	commandIDs []string // we need command ids in order to deregister them later
}

func New(cfg *config.Discord) *Handler {
	return &Handler{
		config: cfg,
		redis:  storage.RedisClient,
		prisma: storage.PrismaClient,
	}
}
