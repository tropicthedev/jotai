package interactionhandler

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) EnrollButton(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	if i.MessageComponentData().CustomID != "action::enroll" {
		return
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: i.Member.AvatarURL("128"),
			Name:    i.Member.User.Username,
		},
		Color:       0x7785cc,                        // Blue
		Timestamp:   time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:       "Beep Boop!",
		Description: h.config.JoinMessage,
	}

	dm, err := s.UserChannelCreate(i.Member.User.ID)

	if err != nil {
		log.Printf("Cannot send message to user '%v, Error: %v", i.Member.User.Username, err)
		return
	}

	s.ChannelMessageSendEmbed(dm.ID, embed)
}
