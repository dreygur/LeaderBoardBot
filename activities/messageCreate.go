package activities

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	update := bson.M{
		"$inc": bson.M{
			"points":          lib.POINTS["text"],
			"activities.text": 1,
		},
	}

	_, err := repo.Collection.UpdateOne(
		context.TODO(),
		bson.M{"user_id": m.Author.ID},
		update,
		options.Update().SetUpsert(true),
	)

	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error in addPointHandler: %v", err), "error")
	}

	lib.PrintLog(fmt.Sprintf("Message from %s: %s", m.Author.Username, m.Content), "info")
}
