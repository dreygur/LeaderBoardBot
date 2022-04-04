package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
)

func GuildInviteCreate(s *discordgo.Session, i *discordgo.InviteCreate) {
	lib.PrintLog(fmt.Sprintf("GuildInviteCreate from %s: %s", i.Inviter.Username, i.Code), "info")
}
