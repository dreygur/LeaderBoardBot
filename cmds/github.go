package cmds

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
)

type repoData struct {
	Items []struct {
		Owner struct {
			Name  string `json:"login"`
			Image string `json:"avatar_url"`
		} `json:"owner"`
		URL      string `json:"html_url"`
		Language string `json:"language"`
		Watchers int    `json:"watchers_count"`
		Forks    int    `json:"forks_count"`
		Issues   int    `json:"open_issues_count"`
		License  struct {
			Name string `json:"name"`
		} `json:"license"`
	} `json:"items"`
}

func getRepoData(repo string) (repoData, error) {
	// Search for Repo
	res, err := http.Get("https://api.github.com/search/repositories?q=" + repo)
	if err != nil {
		return repoData{}, err
	}
	defer res.Body.Close()

	// Read Response
	var data repoData
	json.NewDecoder(res.Body).Decode(&data)
	fmt.Println(data)

	return data, nil
}

func githubHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	repoName := i.ApplicationCommandData().Options[0].StringValue()

	data, err := getRepoData(repoName)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error in github: %v", err), "error")
	}
	message := []*discordgo.MessageEmbed{
		{
			Title:     data.Items[0].Owner.Name,
			Color:     0x00ff00,
			Thumbnail: &discordgo.MessageEmbedThumbnail{URL: data.Items[0].Owner.Image},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Repository",
					Value: data.Items[0].URL,
				},
				{
					Name:  "Most Used Language",
					Value: data.Items[0].Language,
				},
				{
					Name:  "Forks",
					Value: fmt.Sprintf("%d", data.Items[0].Forks),
				},
				{
					Name:   "Watchers",
					Value:  fmt.Sprintf("%d", data.Items[0].Watchers),
					Inline: false,
				},
				{
					Name:  "Open Issues",
					Value: fmt.Sprintf("%d", data.Items[0].Issues),
				},
				{
					Name:  "License",
					Value: data.Items[0].License.Name,
				},
			},
		},
	}

	return message
}

func Github(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: githubHandler(s, i),
		},
	})
}
