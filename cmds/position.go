package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/repo"
)

func positionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	var userName, avatar string
	if len(i.ApplicationCommandData().Options) > 0 {
		userName, avatar = hooks.GetUser(s, i)
	} else {
		userName = i.Member.User.Username
		avatar = i.Member.User.AvatarURL("")
	}

	position, err := repo.Collection.GetPosition(userName)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error getting position: %v", err), "error")
		return []*discordgo.MessageEmbed{
			{
				Title:       "Position",
				Description: "Position of current/specified user not found",
			},
		}
	}

	forAdmin := []*discordgo.MessageEmbed{
		{
			Title:       "Position",
			Description: "Position of current/specified user",
			Author: &discordgo.MessageEmbedAuthor{
				Name: repo.Config.Name,
			},
			Color: 0x3349FF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: avatar,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "User",
					Value:  userName,
					Inline: true,
				},
				{
					Name:   "Position",
					Value:  fmt.Sprint(position),
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
