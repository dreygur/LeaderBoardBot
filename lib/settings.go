package lib

import "github.com/bwmarrin/discordgo"

var BOT_NAME string = "LeaderboardBot"
var LOGO_URL string = "https://i.imgur.com/wSTFkRM.png"
var INVITES = make(map[string][]*discordgo.Invite)
var ACTIVITIES = []string{"text", "reaction", "voice"}
var POINTS = make(map[string]int)

var config = LoadConfig()

func SetPoints() {
	for _, v := range config.Activities {
		POINTS[v] = 5
	}
}

func EditPoints(a string, p int) {
	POINTS[a] = p
}
