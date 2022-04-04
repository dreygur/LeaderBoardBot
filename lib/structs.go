package lib

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name        string
	Description string
	Options     []*discordgo.ApplicationCommandOption
}

type PointPerActivity struct {
	Activity string
	Points   int
}

func (p *PointPerActivity) GetPoints() int {
	return p.Points
}
