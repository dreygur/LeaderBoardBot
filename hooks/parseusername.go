package hooks

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
)

func GetUsername(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	var userID string
	var re = regexp.MustCompile(`(?m)<\@\!(\d.*)>`)

	userName := i.ApplicationCommandData().Options[0].StringValue()

	match := re.FindStringSubmatch(userName)
	if len(match) > 1 {
		userID = match[1]
	}

	member, err := s.State.Member(i.GuildID, userID)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error getting member: %v", err), "error")
	}
	userName = member.User.Username

	return userName
}
