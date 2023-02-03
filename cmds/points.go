package cmds

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/database"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/repo"
)

func getPoint(i *discordgo.InteractionCreate, userName string) (*database.User, error) {
	user, err := repo.Collection.Find(userName)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error getting user: %v", err), "error")
		user, err = repo.CreateIfNotFound(i.Member.User.ID, userName)
		if err != nil {
			errStr := fmt.Sprintf("Error creating user: %v", err)
			lib.PrintLog(errStr, "error")
			return nil, errors.New(errStr)
		}
	}

	return user, nil
}

func pointsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	var userName, avatar string
	if len(i.ApplicationCommandData().Options) > 0 {
		userName, avatar = hooks.GetUser(s, i)
	} else {
		userName = i.Member.User.Username
		avatar = i.Member.User.AvatarURL("")
	}

	forAdmin := []*discordgo.MessageEmbed{
		{
			Title:       "Points",
			Description: "Points of current/specified user",
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
					Name:   "Points",
					Value:  "0",
					Inline: true,
				},
			},
		},
	}

	if hooks.CheckRole(s, i) {
		user, _ := getPoint(i, userName)
		forAdmin[0].Fields[1].Value = fmt.Sprint(user.Points)
		return forAdmin
	}

	if len(i.ApplicationCommandData().Options) > 0 {
		forAdmin[0].Fields[1].Value = "You have to be admin to see other users' points"
	} else {
		user, _ := getPoint(i, userName)
		forAdmin[0].Fields[0].Value = i.Member.User.Username
		forAdmin[0].Fields[1].Value = fmt.Sprint(user.Points)
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
