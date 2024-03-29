package cmds

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/lib"
)

type repoData struct {
	Count int `json:"total_count"`
	Items []struct {
		Owner struct {
			Name  string `json:"login"`
			Image string `json:"avatar_url"`
		} `json:"owner"`
		Name      string `json:"name"`
		URL       string `json:"html_url"`
		Language  string `json:"language"`
		Watchers  int    `json:"watchers_count"`
		Forks     int    `json:"forks_count"`
		Issues    int    `json:"open_issues_count"`
		CreatetAt string `json:"created_at"`
		License   struct {
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

	if data.Count < 1 {
		return repoData{}, fmt.Errorf("no repo found")
	}

	return data, nil
}

func githubHandler(s *discordgo.Session, i *discordgo.InteractionCreate) []*discordgo.MessageEmbed {
	repoName := i.ApplicationCommandData().Options[0].StringValue()

	data, err := getRepoData(repoName)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error in github: %v", err), "error")
		return []*discordgo.MessageEmbed{
			{
				Title:       "Error",
				Description: fmt.Sprint(err, " named ", repoName),
				Color:       0x4682B4,
			},
		}
	}

	parsedTime, err := time.Parse(time.RFC3339, data.Items[0].CreatetAt)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error in github: %v", err), "error")
	}

	message := []*discordgo.MessageEmbed{
		{
			Color: 0x00ff00,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    data.Items[0].Owner.Name,
				IconURL: data.Items[0].Owner.Image,
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{URL: data.Items[0].Owner.Image},
			Fields:    []*discordgo.MessageEmbedField{},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Repo created at " + parsedTime.Format("02/01/2006"),
			},
		},
	}

	// Ripository Field
	if data.Items[0].Name != "" && data.Items[0].URL != "" {
		message[0].Fields = append(message[0].Fields, &discordgo.MessageEmbedField{
			Name:   "Repository",
			Value:  fmt.Sprintf("[%s](%s)", data.Items[0].Name, data.Items[0].URL),
			Inline: true,
		})
	}

	// Most Used Language Field
	if data.Items[0].Language != "" {
		message[0].Fields = append(message[0].Fields, &discordgo.MessageEmbedField{
			Name:   "Most Used Language",
			Value:  data.Items[0].Language,
			Inline: true,
		})
	}

	// Forks Field
	if data.Items[0].Forks != 0 {
		message[0].Fields = append(message[0].Fields, &discordgo.MessageEmbedField{
			Name:   "Forks",
			Value:  fmt.Sprintf("%d", data.Items[0].Forks),
			Inline: true,
		})
	}

	// Watchers Field
	if data.Items[0].Watchers != 0 {
		message[0].Fields = append(message[0].Fields, &discordgo.MessageEmbedField{
			Name:   "Watchers",
			Value:  fmt.Sprintf("%d", data.Items[0].Watchers),
			Inline: true,
		})
	}

	// Open Issues Field
	if data.Items[0].Issues != 0 {
		message[0].Fields = append(message[0].Fields, &discordgo.MessageEmbedField{
			Name:   "Open Issues",
			Value:  fmt.Sprintf("%d", data.Items[0].Issues),
			Inline: true,
		})
	}

	// License Field
	if data.Items[0].License.Name != "" {
		message[0].Fields = append(message[0].Fields, &discordgo.MessageEmbedField{
			Name:   "License",
			Value:  data.Items[0].License.Name,
			Inline: true,
		})
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
