package cmds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/repo"
)

func pointsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	var userName string
	if len(i.ApplicationCommandData().Options) > 0 {
		userName = hooks.GetUsername(s, i)
	} else {
		userName = i.Member.User.Username
	}

	// var user *database.User
	// err := repo.Collection.FindOne(context.Background(), bson.M{"username": userName}).Decode(&user)
	user, err := repo.Collection.Find(userName)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error getting user: %v", err), "error")
	}

	forAdmin := []*discordgo.MessageEmbed{
		{
			Title:       "Points",
			Description: "Points",
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
					Name:   "Points",
					Value:  fmt.Sprint(user.Points),
					Inline: true,
				},
			},
		},
	}

	if hooks.CheckRole(s, i) {
		return forAdmin
	}

	if len(i.ApplicationCommandData().Options) > 0 {
		forAdmin[0].Fields[1].Value = "You have to be admin to see other users' points"
	} else {
		forAdmin[0].Fields[0].Value = i.Member.User.Username
	}

	return forAdmin
}

func Points(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: pointsHandler(s, i),
		},
	})
}
