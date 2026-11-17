package hooks

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/mention"
)

func GetUser(s *discordgo.Session, i *discordgo.InteractionCreate) (string, string) {
	userID := mention.UserID(i.ApplicationCommandData().Options[0].StringValue())

	member, err := s.State.Member(i.GuildID, userID)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error getting member: %v", err), "error")
	}
	userName := member.User.Username
	profileImage := member.User.AvatarURL("")

	return userName, profileImage
}
