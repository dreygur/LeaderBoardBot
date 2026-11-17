package db

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/dreygur/leaderboardbot/database"
	"github.com/dreygur/leaderboardbot/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (d *MongoDB) CreateNewUser(u database.User) error {
	res, err := d.Instance.InsertOne(context.Background(), u)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Failed to insert user into database %v", err), "error")
		return err
	}
	lib.PrintLog(fmt.Sprintf("Inserted user %s into database with ID %v", u.Username, res.InsertedID), "info")

	return nil
}

func (d *MongoDB) GetPosition(userName string) (int, error) {
	cursor, err := d.Instance.Find(context.TODO(), bson.M{})
	if err != nil {
		return 0, err
	}

	var users []*database.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Points > users[j].Points
	})

	for i, user := range users {
		if user.Username == userName {
			return i + 1, nil
		}
	}

	return 0, fmt.Errorf("user %s not found", userName)
}

func (d *MongoDB) Update(target, data bson.M) error {
	_, err := d.Instance.UpdateOne(
		context.TODO(),
		target,
		data,
		options.UpdateOne().SetUpsert(true),
	)

	if err != nil {
		return err
	}

	return nil
}

// Increment applies data to the user matching target, creating the user if
// missing, and returns the document after the change so callers can act on
// the new totals without a second read.
func (d *MongoDB) Increment(target, data bson.M) (*database.User, error) {
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var user database.User
	err := d.Instance.FindOneAndUpdate(context.TODO(), target, data, opts).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("increment user %v: %w", target, err)
	}

	return &user, nil
}

// Register creates the user unless one with userID already exists, and
// reports whether it did. Records outlive leaving the server, so this is
// true only on a user's first ever join.
func (d *MongoDB) Register(userID, userName string) (bool, error) {
	res, err := d.Instance.UpdateOne(
		context.TODO(),
		bson.M{"user_id": userID},
		bson.M{"$setOnInsert": bson.M{"username": userName}},
		options.UpdateOne().SetUpsert(true),
	)
	if err != nil {
		return false, fmt.Errorf("register user %s: %w", userID, err)
	}

	return res.UpsertedCount == 1, nil
}

func (d *MongoDB) Delete() error {
	return nil
}

func (d *MongoDB) Find(userName string) (*database.User, error) {
	var user *database.User
	err := d.Instance.FindOne(context.TODO(), bson.M{"username": userName}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
