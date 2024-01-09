package interactionhandler

import (
	"context"
	"fmt"
	"jordanmckoy/guardian/internal/storage"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) DirectMessage(s *discordgo.Session, i *discordgo.MessageCreate) {
	var ctx = context.Background()

	if i.Author.Bot {
		return
	}

	if i.GuildID == "" {
		fmt.Println(storage.GetClient().Get(ctx, i.Author.ID).Result())
	}

}
