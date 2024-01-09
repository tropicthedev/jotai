package interactionhandler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h Handler) TearUpCommands(session *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Ping the bot. Just for fun.",
		},
		{
			Name:        "enroll",
			Description: "Sends a embed with a link to full out the server application.",
		},

		// {
		// 	Name:        "accept",
		// 	Description: "Accept a user into the server.",
		// },
		// {
		// 	Name:        "reject",
		// 	Description: "Reject a user from the server.",
		// },
		// {
		// 	Name:        "whitelist-add",
		// 	Description: "Add a user to the whitelist.",
		// },
		// {
		// 	Name:        "whitelist-remove",
		// 	Description: "Remove a user from the whitelist.",
		// },
		// {
		// 	Name:        "audit",
		// 	Description: "Audits the servers whitelist and removes users who have not been active for a certain amount of days.",
		// },
	}

	for _, c := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, h.config.GuildID, c)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", cmd.Name, err)
		}

		h.commandIDs = append(h.commandIDs, cmd.ID)
	}
}

func (h Handler) TearDownCommands(session *discordgo.Session) {
	for _, cmd := range h.commandIDs {
		err := session.ApplicationCommandDelete(session.State.User.ID, h.config.GuildID, cmd)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", cmd, err)
		}
	}
}
