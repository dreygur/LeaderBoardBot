package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/repo"
)

func positionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	var userName string
	if len(i.ApplicationCommandData().Options) > 0 {
		userName = hooks.GetUsername(s, i)
	} else {
		userName = i.Member.User.Username
	}

	forAdmin := []*discordgo.MessageEmbed{
		{
			Title:       "Position",
			Description: "Position",
			Author: &discordgo.MessageEmbedAuthor{
				Name: repo.Config.Name,
			},
			Color: 0x3349FF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: repo.Config.LogoURL,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "User",
					Value:  userName,
					Inline: true,
				},
				{
					Name:   "Position",
					Value:  "5",
					Inline: true,
				},
			},
		},
	}

	if hooks.CheckRole(s, i) {
		return forAdmin
	}

	if len(i.ApplicationCommandData().Options) > 0 {
		forAdmin[0].Fields[1].Value = "You have to be admin to get points"
	} else {
		forAdmin[0].Fields[0].Value = i.Member.User.Username
		forAdmin[0].Fields[1].Value = "0"
	}

	return forAdmin
}

func Position(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: positionHandler(s, i),
		},
	})
}
