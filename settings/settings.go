package settings

import (
	"github.com/dreygur/leaderboardbot/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config interface {
	SetDB(db *database.Database) error
	GetCollection() *mongo.Collection
}

type Conf struct{}

func (c *Conf) SetDB(db *database.Database) *Config {
	return nil
}

func (c *Conf) GetCollection() *Config {
	return nil
}
