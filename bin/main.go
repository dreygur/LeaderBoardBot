package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/activities"
	"github.com/dreygur/leaderboardbot/database"
	"github.com/dreygur/leaderboardbot/events"
	"github.com/dreygur/leaderboardbot/handlers"
	"github.com/dreygur/leaderboardbot/lib"
)

func main() {
	// godotenv.Load()
	config := lib.LoadConfig()
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(err)
	}

	// Select the intents
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	// dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Detect Login Event
	dg.AddHandler(events.LoginEvent)
	// New User Join Event
	dg.AddHandler(events.GuildMemberAdd)
	// Command Handler
	dg.AddHandler(events.InteractionCreate)
	// Invite Create Handler
	dg.AddHandler(events.GuildInviteCreate)
	// New Message Event
	dg.AddHandler(activities.MessageCreate)
	// Reaction Create Handler
	dg.AddHandler(activities.ReactionAdd)
	// Voice State Update Handler
	dg.AddHandler(activities.VoiceStateUpdate)

	// Start the bot
	err = dg.Open()
	if err != nil {
		panic(err)
	}
	// Wait here until CTRL-C or other term signal is received.
	lib.PrintLog("Bot is now running.  Press CTRL-C to exit.", "info")

	// Register Commands
	handlers.InitCommands(dg)
	// Remove Commands
	// command.RemoveCommands(dg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sig

	// Stop the bot
	defer dg.Close()
	defer database.Disconnect()
}
