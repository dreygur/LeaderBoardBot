package repo

import (
	"github.com/dreygur/leaderboardbot/database"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/settings"
)

var (
	// Configurations
	Config = lib.LoadConfig()

	// Database Connection
	Collection = settings.NewDatabase(database.Database{
		Address:    Config.DatabaseURL,
		Name:       Config.Database.Name,
		Collection: Config.Database.Collection,
	}).GetCollection()
)
