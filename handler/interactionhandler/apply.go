package interactionhandler

import (
	"github.com/bwmarrin/discordgo"
)

func (h Handler) Apply(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name != "apply" {
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Check your DMs!",
		},
	})
}
