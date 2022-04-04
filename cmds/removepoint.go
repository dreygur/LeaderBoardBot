package cmds

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func removePointHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	defer func() {
		if r := recover(); r != nil {
			lib.PrintLog(fmt.Sprintf("Recovered in addPointHandler: %v", r), "error")
		}
	}()

	userName := hooks.GetUsername(s, i)
	points := i.ApplicationCommandData().Options[1].IntValue()

	forAdmin := []*discordgo.MessageEmbed{
		{
			Title:       "Remove Point",
			Description: "Remove point from user",
			Author: &discordgo.MessageEmbedAuthor{
				Name: config.Name,
			},
			Color: 0xD4122C,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: config.LogoURL,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Points",
					Value: fmt.Sprintf("%d points removed from user %s", points, userName),
				},
			},
		},
	}

	forUser := []*discordgo.MessageEmbed{
		{
			Title:       "Remove Point",
			Description: "Remove point from user",
			Author: &discordgo.MessageEmbedAuthor{
				Name: config.Name,
			},
			Color: 0xD4122C,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: config.LogoURL,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Points",
					Value: "You have to be admin to remove points",
				},
			},
		},
	}

	if hooks.CheckRole(s, i) {
		_, err := collection.UpdateOne(context.TODO(), bson.M{"username": userName}, bson.M{"$inc": bson.M{"points": -points}}, options.Update().SetUpsert(true))
		if err != nil {
			lib.PrintLog(fmt.Sprintf("Error in addPointHandler: %v", err), "error")
		}

		lib.PrintLog(fmt.Sprintf("Updated user: %s", userName), "info")
		return forAdmin
	}

	return forUser
}

func RemovePoint(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: removePointHandler(s, i),
		},
	})
}
