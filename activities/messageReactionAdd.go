package activities

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReactionAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.UserID == s.State.User.ID {
		return
	}

	update := bson.M{
		"$inc": bson.M{
			"points":              lib.POINTS["reaction"],
			"activities.reaction": lib.POINTS["reaction"],
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"user_id": m.UserID}, update, options.Update().SetUpsert(true))

	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error in addPointHandler: %v", err), "error")
	}

	lib.PrintLog(fmt.Sprintf("Reaction from %s: %s", m.Member.User.Username, m.Emoji.Name), "info")
}
