package db

import (
	"github.com/dreygur/leaderboardbot/database"
	"github.com/dreygur/leaderboardbot/lib"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	// Configurations
	Config = lib.LoadConfig()
	err    error
)

// Database Interface
type Database interface {
	Connect() error
	CreateNewUser(u database.User) error
	Find(userName string) (*database.User, error)
	GetPosition(userName string) (int, error)
	Update(target, data bson.M) error
	Increment(target, data bson.M) (*database.User, error)
	Delete() error
	Close() error
}

// MongoDB Database
func GetDatabase(d Database) *MongoDB {
	return &MongoDB{
		Address:    Config.DatabaseURL,
		Name:       Config.Database.Name,
		Collection: Config.Database.Collection,
	}
}

type MongoDB struct {
	Address    string            `json:"db_url"`
	Name       string            `json:"name"`
	Collection string            `json:"collection"`
	Instance   *mongo.Collection // MongoDB Collection Instance
	Client     *mongo.Client     // mongo client
}
