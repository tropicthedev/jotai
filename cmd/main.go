package main

import (
	"flag"
	"jordanmckoy/guardian/config"
	"jordanmckoy/guardian/delivery/websocketserver"
	"jordanmckoy/guardian/handler/interactionhandler"
	"jordanmckoy/guardian/handler/messagehandler"
	"jordanmckoy/guardian/internal/storage"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	cfg config.Config
)

func init() {
	cfg = config.Load("config.yml")
	flag.Parse()
}

func main() {
	// add as many handlers as you want implementing websocketserver.Handler...
	handlers := []websocketserver.Handler{
		messagehandler.New(&cfg.Discord),
		interactionhandler.New(&cfg.Discord),
	}

	// Redis
	storage.RedisInit()

	// Prisma
	storage.PrismaInit()

	server := websocketserver.New(&cfg, handlers, discordgo.IntentsAll)
	server.Serve()
	defer server.Shutdown()

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
