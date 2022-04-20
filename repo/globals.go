package repo

import (
	"github.com/dreygur/leaderboardbot/db"
	"github.com/dreygur/leaderboardbot/lib"
)

var (
	// Configurations
	Config = lib.LoadConfig()

	// Database Connection
	Collection = db.GetDatabase(&db.MongoDB{
		Address:    Config.DatabaseURL,
		Name:       Config.Database.Name,
		Collection: Config.Database.Collection,
	})
)
