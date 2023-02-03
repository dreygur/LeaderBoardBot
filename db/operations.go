package db

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/dreygur/leaderboardbot/database"
	"github.com/dreygur/leaderboardbot/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (d *MongoDB) Update(target, data primitive.M) error {
	_, err := d.Instance.UpdateOne(
		context.TODO(),
		target,
		data,
		options.Update().SetUpsert(true),
	)

	if err != nil {
		return err
	}

	return nil
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
