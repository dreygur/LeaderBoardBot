package db

import (
	"context"
	"fmt"

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

func (d *MongoDB) Add() error {
	// res, err := d.Instance.InsertOne(context.Background(), database.User{
	// 	targets[0]: m.User.ID,
	// 	Username:   m.User.Username,
	// 	Points:     0,
	// 	Level:      0,
	// 	Activities: database.Activity{
	// 		Text:     0,
	// 		Reaction: 0,
	// 		Voice:    0,
	// 	},
	// })
	// if err != nil {
	// 	lib.PrintLog(fmt.Sprintf("Failed to insert user into database %v", err), "error")
	// 	return err
	// }
	// lib.PrintLog(fmt.Sprintf("Inserted user %s into database with ID %v", m.User.Username, res.InsertedID), "info")

	return nil
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
