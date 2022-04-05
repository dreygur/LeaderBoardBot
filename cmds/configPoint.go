package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/database"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/settings"
)

var (
	// Configurations
	config = lib.LoadConfig()

	// Database Connection
	collection = settings.NewDatabase(database.Database{
		Address:    config.DatabaseURL,
		Name:       config.Database.Name,
		Collection: config.Database.Collection,
	}).GetCollection()
)

func configPointHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	activity := i.ApplicationCommandData().Options[0].StringValue()
	point := i.ApplicationCommandData().Options[1].IntValue()

	lib.POINTS[activity] = int(point)

	forAdmin := []*discordgo.MessageEmbed{
		{
			Title:       "Points",
			Description: "Points",
			Author: &discordgo.MessageEmbedAuthor{
				Name: config.Name,
			},
			Color: 0x3349FF,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: config.LogoURL,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Activity",
					Value:  activity,
					Inline: true,
				},
				{
					Name:   "Points",
					Value:  fmt.Sprintf("%d", point),
					Inline: true,
				},
			},
		},
	}

	if hooks.CheckRole(s, i) {
		return forAdmin
	}

	forAdmin[0].Fields[1].Value = "You have to be admin to add points"

	return forAdmin
}

func ConfigPoint(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: configPointHandler(s, i),
		},
	})
}
