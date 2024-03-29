package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/repo"
	"go.mongodb.org/mongo-driver/bson"
)

func addPointHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	defer func() {
		if r := recover(); r != nil {
			lib.PrintLog(fmt.Sprintf("Recovered in addPointHandler: %v", r), "error")
		}
	}()

	userName, avatar := hooks.GetUser(s, i)
	points := i.ApplicationCommandData().Options[1].IntValue()

	forAdmin := []*discordgo.MessageEmbed{
		{
			Title:       "Add Point",
			Description: "Add point to user",
			Author: &discordgo.MessageEmbedAuthor{
				Name: repo.Config.Name,
			},
			Color: 0x179ED1,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: avatar,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Points",
					Value: fmt.Sprintf("%d points added to user %s", points, userName),
				},
			},
		},
	}

	forUser := []*discordgo.MessageEmbed{
		{
			Title:       "Add Point",
			Description: "Add point to user",
			Author: &discordgo.MessageEmbedAuthor{
				Name: repo.Config.Name,
			},
			Color: 0x179ED1,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: repo.Config.LogoURL,
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Points",
					Value: "You have to be admin to add points",
				},
			},
		},
	}

	if hooks.CheckRole(s, i) {
		update := bson.M{
			"$inc": bson.M{
				"points": int(points),
			},
		}
		err := repo.Collection.Update(
			bson.M{"user_id": i.Member.User.ID},
			update,
		)

		if err != nil {
			lib.PrintLog(fmt.Sprintf("Error in addPointHandler: %v", err), "error")
		}

		lib.PrintLog(fmt.Sprintf("Updated user: %s", userName), "info")
		return forAdmin
	}

	return forUser
}

func AddPoint(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: addPointHandler(s, i),
		},
	})
}
