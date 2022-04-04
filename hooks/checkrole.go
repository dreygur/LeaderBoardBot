package hooks

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
)

func CheckRole(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	config := lib.LoadConfig()
	for _, v := range config.Roles {
		for _, r := range i.Member.Roles {
			userRole, err := s.State.Role(i.GuildID, r)
			if err != nil {
				fmt.Println(err)
			}
			if strings.EqualFold(userRole.Name, v) {
				return true
			}
		}
	}
	return false
}
