package activities

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
)

func VoiceStateUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	if v.UserID == s.State.User.ID {
		return
	}

	user, err := s.State.Member(v.GuildID, v.UserID)
	if err != nil {
		lib.PrintLog("Failed to get user name for "+v.UserID, "error")
		return
	}
	lib.PrintLog(fmt.Sprintf("VoiceStateUpdate from %s: %s", user.User.Username, v.ChannelID), "info")
}
