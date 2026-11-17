package repo

import (
	"fmt"

	"github.com/dreygur/leaderboardbot/scoring"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Award applies inc as a $inc to the user, creating them if missing, then
// brings their stored level in line with their new points. It returns the
// new level when the user went up, or 0. A drop in level is stored silently.
func Award(userID string, inc bson.M) (int, error) {
	filter := bson.M{"user_id": userID}

	user, err := Collection.Increment(filter, bson.M{"$inc": inc})
	if err != nil {
		return 0, err
	}

	level := scoring.LevelFor(user.Points)
	if level == user.Level {
		return 0, nil
	}

	err = Collection.Update(filter, bson.M{"$set": bson.M{"level": level}})
	if err != nil {
		return 0, fmt.Errorf("set level %d for %s: %w", level, userID, err)
	}

	if level < user.Level {
		return 0, nil
	}
	return level, nil
}
