package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/activities"
	"github.com/dreygur/leaderboardbot/events"
	"github.com/dreygur/leaderboardbot/handlers"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/repo"
)

func main() {
	// Recover From Panicing
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
		}
	}()

	// godotenv.Load()
	config := lib.LoadConfig()

	// Initiate Database Connection First
	err := repo.Collection.Connect()
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error connecting Database: %v", err), "error")
		return
	}
	defer repo.Collection.Close()

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

	/**
	 * Not Disconnecting Database this time
	 * Will implement this one in future
	 * though this is an intended bug!
	 */
	// defer database.Disconnect()
}
