package lib

import (
	"github.com/bwmarrin/discordgo"
)

// RegisteredCommands is a slice of ApplicationCommand
var RegisteredCommands []*discordgo.ApplicationCommand

// Commands
var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "help",
		Description: "Shows help message",
	},
	{
		Name:        "addpoint",
		Description: "Add point to user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Enter user's name",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "points",
				Description: "Enter points",
				Required:    true,
			},
		},
	},
	{
		Name:        "removepoint",
		Description: "Remove point from user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Enter activity name",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "points",
				Description: "Enter points",
				Required:    true,
			},
		},
	},
	{
		Name:        "configpoints",
		Description: "Configure points",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Enter activity name",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "points",
				Description: "Enter points",
				Required:    true,
			},
		},
	},
	{
		Name:        "points",
		Description: "Show points",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Enter target user's name or leave blank to see your points",
				Required:    false,
			},
		},
	},
	{
		Name:        "position",
		Description: "Show position",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Enter target user's name or leave blank to see your position",
				Required:    false,
			},
		},
	},
	{
		Name:        "userinvited",
		Description: "Show user invited count",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Enter target user's name",
				Required:    true,
			},
		},
	},
	{
		Name:        "leaderboard",
		Description: "Show leaderboard",
	},
	{
		Name:        "play",
		Description: "Play a music from url",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "url",
				Description: "Enter the URL to play music",
				Required:    true,
			},
		},
	},
	{
		Name:        "stop",
		Description: "Stops playing music",
	},
	{
		Name:        "pause",
		Description: "Pauses playing music",
	},
}
