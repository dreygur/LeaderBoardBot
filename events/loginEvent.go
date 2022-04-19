package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
)

func LoginEvent(s *discordgo.Session, r *discordgo.Ready) {
	// Log the new Username to Console
	lib.PrintLog("Logged in as "+r.User.Username, "info")

	// Set points
	lib.SetPoints()

	for _, v := range r.Guilds {
		invites, err := s.GuildInvites(v.ID)
		if err != nil {
			lib.PrintLog("Failed to get invites for "+v.ID, "error")
			return
		}
		lib.INVITES[v.ID] = invites
	}
}
