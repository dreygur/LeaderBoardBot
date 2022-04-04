package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/handlers"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if function, ok := handlers.CommandHandlers[i.ApplicationCommandData().Name]; ok {
		function(s, i)
	}
}
