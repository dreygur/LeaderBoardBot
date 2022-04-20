package db

import (
	"context"
	"fmt"
	"time"

	"github.com/dreygur/leaderboardbot/lib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *MongoDB) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	d.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(d.Address))
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error connecting to database: %v", err), "error")
		return err
	}

	if d.Client.Ping(ctx, nil) != nil {
		lib.PrintLog(fmt.Sprintf("Error pinging database: %v", err), "error")
		return fmt.Errorf("error pinging database: %v", err)
	}

	d.Instance = d.Client.Database(d.Name).Collection(d.Collection)
	return nil
}

func (d *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = d.Client.Disconnect(ctx)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error disconnecting from database: %v", err), "error")
		return err
	}

	return nil
}
