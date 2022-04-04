package hooks

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
)

func CheckInviter(m *discordgo.GuildMemberAdd, invites []*discordgo.Invite) string {
	for _, v := range lib.INVITES[m.GuildID] {
		for _, r := range invites {
			if strings.EqualFold(r.Inviter.Username, v.Inviter.Username) && r.Uses != v.Uses {
				return r.Inviter.Username
			}
		}
	}

	return ""
}
