package database

import (
	"context"
	"fmt"
	"time"

	"github.com/dreygur/leaderboardbot/lib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

func ConnectDB() *mongo.Collection {
	config := lib.LoadConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.DatabaseURL))
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error connecting to database: %v", err), "error")
	}

	if client.Ping(ctx, nil) != nil {
		lib.PrintLog(fmt.Sprintf("Error pinging database: %v", err), "error")
	}

	database := client.Database("leaderboard")
	userCollection := database.Collection("users")
	return userCollection
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Disconnect(ctx)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error disconnecting from database: %v", err), "error")
	}
}
