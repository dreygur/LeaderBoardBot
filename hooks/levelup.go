package hooks

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
)

const levelUpColor = 0xF1C40F

// AnnounceLevel posts a level up embed for userID in channelID.
func AnnounceLevel(s *discordgo.Session, channelID, userID string, level int) {
	_, err := s.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Title:       "Level Up",
		Description: fmt.Sprintf("<@%s> reached level %d!", userID, level),
		Color:       levelUpColor,
	})
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Failed to announce level %d for %s: %v", level, userID, err), "error")
	}
}

// DMLevel sends the level up embed privately, for activity that has no
// text channel to answer in, like voice or invites.
func DMLevel(s *discordgo.Session, userID string, level int) {
	dm, err := s.UserChannelCreate(userID)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Failed to open DM with %s: %v", userID, err), "error")
		return
	}
	AnnounceLevel(s, dm.ID, userID, level)
}
