package database

import (
	"context"
	"fmt"
	"time"

	"github.com/dreygur/leaderboardbot/lib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var err error

type Database struct {
	Address    string
	Name       string
	Collection string
	Client     *mongo.Client
}

func (d *Database) ConnectDB() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	d.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(d.Address))
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error connecting to database: %v", err), "error")
	}

	if d.Client.Ping(ctx, nil) != nil {
		lib.PrintLog(fmt.Sprintf("Error pinging database: %v", err), "error")
	}

	database := d.Client.Database(d.Name)
	userCollection := database.Collection(d.Collection)
	return userCollection
}

func (d *Database) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = d.Client.Disconnect(ctx)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error disconnecting from database: %v", err), "error")
	}
}
