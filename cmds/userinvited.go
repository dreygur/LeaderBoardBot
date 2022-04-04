package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/hooks"
)

func userInvitedHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
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
				Name: config.Name,
			},
			Color: 0x1232D4,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: config.LogoURL,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "User",
					Value:  userName,
					Inline: true,
				},
				{
					Name:   "Count",
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
		forAdmin[0].Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "User",
				Value: "You have to be admin to execute this command",
			},
		}
	}

	return forAdmin
}

func UserInvited(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: userInvitedHandler(s, i),
		},
	})
}
