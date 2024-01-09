package interactionhandler

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) DirectMessage(s *discordgo.Session, i *discordgo.MessageCreate) {
	if i.Author.Bot {
		return
	}

	if i.GuildID != "" {
		return
	}

	var ctx = context.Background()

	val := h.redis.Get(ctx, i.Author.ID).Val()

	if val != "started" {
		return
	}

	log.Println(i.Message.Content)
}
