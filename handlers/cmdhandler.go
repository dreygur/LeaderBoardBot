package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/cmds"
	"github.com/dreygur/leaderboardbot/lib"
)

func InitCommands(s *discordgo.Session) {
	lib.PrintLog("Adding commands...", "info")
	lib.RegisteredCommands = make([]*discordgo.ApplicationCommand, len(lib.Commands))

	for i, v := range lib.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", cmd.Name, err)
		}
		lib.RegisteredCommands[i] = cmd
		lib.PrintLog("Added command: "+cmd.Name, "info")
	}
}

func RemoveCommands(s *discordgo.Session) {
	lib.PrintLog("Removing commands...", "info")
	for _, v := range lib.RegisteredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
		lib.PrintLog("Removed command: "+v.Name, "info")
	}
}

var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"help":         cmds.Help,
	"addpoint":     cmds.AddPoint,
	"points":       cmds.Points,
	"removepoint":  cmds.RemovePoint,
	"position":     cmds.Position,
	"userinvited":  cmds.UserInvited,
	"configpoints": cmds.ConfigPoint,
	"play":         cmds.GetMusic,
	"stop":         cmds.StopMusic,
	"pause":        cmds.PauseMusic,
	"github":       cmds.Github,
}
