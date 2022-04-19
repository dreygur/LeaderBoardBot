package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/repo"
)

func HelpMessageHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	forAdmin := []*discordgo.MessageEmbed{
		{
			Title:       "Help",
			Description: "This is a help message",
			Author: &discordgo.MessageEmbedAuthor{
				Name: repo.Config.Name,
			},
			Color: 0x00ff00,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: repo.Config.LogoURL,
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    "https://i.imgur.com/w3duR07.png",
				Height: 300,
				Width:  300,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "`/help`",
					Value:  "Shows this help message",
					Inline: true,
				},
				{
					Name:   "`/addpoint`",
					Value:  "Adds point to user",
					Inline: false,
				},
				{
					Name:   "`/removepoint`",
					Value:  "Remove point from user",
					Inline: false,
				},
				{
					Name:   "`/configpoints`",
					Value:  "Configure points",
					Inline: false,
				},
				{
					Name:  "`/leaderboard`",
					Value: "Shows leaderboard",
				},
				{
					Name:  "`/points`",
					Value: "Shows your points",
				},
				{
					Name:  "`/position`",
					Value: "Shows your position on LeaderBoard",
				},
				{
					Name:  "`/userinvited`",
					Value: "Shows how many users you have invited",
				},
				{
					Name:  "`/play`",
					Value: "Play a music from yt!",
				},
				{
					Name:  "`/stop`",
					Value: "Stop playing music",
				},
				{
					Name:  "`/pause`",
					Value: "Pause currently playing music",
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: repo.Config.Name,
			},
		},
	}

	forUsers := []*discordgo.MessageEmbed{
		{
			Title:       "Help",
			Description: "This is a help message",
			Author: &discordgo.MessageEmbedAuthor{
				Name: repo.Config.Name,
			},
			Color: 0x00ff00,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: repo.Config.LogoURL,
			},
			Image: &discordgo.MessageEmbedImage{
				URL:    "https://i.imgur.com/w3duR07.png",
				Height: 300,
				Width:  300,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "`/help`",
					Value:  "Shows this help message",
					Inline: true,
				},
				{
					Name:  "`/points`",
					Value: "Shows your points",
				},
				{
					Name:  "`/position`",
					Value: "Shows your position on LeaderBoard",
				},
				{
					Name:  "`/userinvited`",
					Value: "Shows how many users you have invited",
				},
				{
					Name:  "`/play`",
					Value: "Play a music from yt!",
				},
				{
					Name:  "`/stop`",
					Value: "Stop playing music",
				},
				{
					Name:  "`/pause`",
					Value: "Pause currently playing music",
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: repo.Config.Name,
			},
		},
	}
	if hooks.CheckRole(s, i) {
		return forAdmin
	}
	return forUsers
}

func Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: HelpMessageHandler(s, i),
		},
	})
}
