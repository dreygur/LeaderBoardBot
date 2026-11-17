package activities

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/repo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ReactionAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.UserID == s.State.User.ID {
		return
	}

	update := bson.M{
		"points":              lib.POINTS["reaction"],
		"activities.reaction": 1,
	}

	// _, err := repo.Collection.UpdateOne(
	// 	context.TODO(),
	// 	bson.M{"user_id": m.UserID},
	// 	update,
	// 	options.Update().SetUpsert(true),
	// )

	level, err := repo.Award(m.UserID, update)

	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error in addPointHandler: %v", err), "error")
	}
	if level > 0 {
		hooks.AnnounceLevel(s, m.ChannelID, m.UserID, level)
	}

	lib.PrintLog(fmt.Sprintf("Reaction from %s: %s", m.Member.User.Username, m.Emoji.Name), "info")
}
