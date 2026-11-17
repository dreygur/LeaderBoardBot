package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/database"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/repo"
	"github.com/dreygur/leaderboardbot/scoring"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func awardInvite(s *discordgo.Session, inviterID string) {
	level, err := repo.Award(inviterID, bson.M{
		"points":  lib.POINTS["invite"],
		"invites": 1,
	})
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error awarding invite points: %v", err), "error")
	}
	if level > 0 {
		hooks.DMLevel(s, inviterID, level)
	}
}

func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Log the new Username to Console
	lib.PrintLog(m.User.Username+" has joined the server.", "info")

	server, err := s.Guild(m.GuildID)
	if err != nil {
		lib.PrintLog("Failed to get guild info for "+m.GuildID, "error")
		return
	}

	// Invites run before the DM so a member with closed DMs still credits
	// their inviter and keeps the cached use counts current.
	invites, err := s.GuildInvites(m.GuildID)
	if err != nil {
		lib.PrintLog("Failed to get invites for "+m.GuildID, "error")
		return
	}

	inviter := scoring.Invites.Used(m.GuildID, invites)
	if inviter != nil {
		lib.PrintLog(fmt.Sprintf("Inviter of %s is %s", m.User.Username, inviter.Username), "info")
		awardInvite(s, inviter.ID)
	}

	// Create a Private Channel with the user to send DM
	dm, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		lib.PrintLog("Failed to create DM channel for "+m.User.Username, "error")
		return
	}

	// Save the user to the database
	repo.Collection.CreateNewUser(database.User{
		UserID:   m.User.ID,
		Username: m.User.Username,
		Points:   0,
		Level:    0,
		Activities: database.Activity{
			Text:     0,
			Reaction: 0,
			Voice:    0,
		},
	})

	// Save to Database
	// res, err := repo.Collection.InsertOne(context.TODO(), database.User{
	// 	UserID:   m.User.ID,
	// 	Username: m.User.Username,
	// 	Points:   0,
	// 	Level:    0,
	// 	Activities: database.Activity{
	// 		Text:     0,
	// 		Reaction: 0,
	// 		Voice:    0,
	// 	},
	// })
	// if err != nil {
	// 	lib.PrintLog(fmt.Sprintf("Failed to insert user into database %v", err), "error")
	// }
	// lib.PrintLog(fmt.Sprintf("Inserted user %s into database with ID %v", m.User.Username, res.InsertedID), "info")

	// Send DM to the new user
	// s.ChannelMessageSend(dm.ID, "Hello "+m.User.Username+"!")
	s.ChannelMessageSendEmbed(dm.ID, &discordgo.MessageEmbed{
		Title:       "Welcome to " + server.Name,
		Description: "Hi " + m.User.Username + "!\nPlease read the rules and have fun!\n",
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Leaderboard Bot",
			IconURL: lib.LOGO_URL,
		},
		Color: 0x00ff00,
	})
}
