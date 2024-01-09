package interactionhandler

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) Enroll(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if i.ApplicationCommandData().Name != "enroll" {
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

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "📜",
							},
							Label:    "Apply",
							Style:    discordgo.PrimaryButton,
							CustomID: "action::enroll",
						},
					},
				},
			},
		},
	})
}
