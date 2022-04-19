package settings

import (
	"github.com/dreygur/leaderboardbot/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config interface {
	GetCollection() *mongo.Collection
	Diconnect()
}

var (
	d database.Database
)

type conf struct{}

func NewDatabase(db database.Database) Config {
	d = db
	return &conf{}
}

func (*conf) GetCollection() *mongo.Collection {
	return d.ConnectDB()
}

func (*conf) Diconnect() {
	d.Disconnect()
}
