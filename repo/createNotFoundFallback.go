package repo

import (
	"fmt"

	"github.com/dreygur/leaderboardbot/database"
	"github.com/dreygur/leaderboardbot/lib"
)

func CreateIfNotFound(userId, userName string) (*database.User, error) {
	err := Collection.CreateNewUser(database.User{
		UserID:   userId,
		Username: userName,
		Points:   0,
		Level:    0,
		Activities: database.Activity{
			Text:     0,
			Reaction: 0,
			Voice:    0,
		},
	})
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error creating user: %v", err), "error")
		return nil, err
	}
	lib.PrintLog(fmt.Sprintf("Created user: %s", userName), "info")
	user, err := Collection.Find(userName)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error finding user: %v", err), "error")
		return nil, err
	}

	return user, nil
}
