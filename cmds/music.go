package cmds

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	ytdl "github.com/kkdai/youtube/v2"
)

type Youtube struct {
	Title   string `json:"title"`
	Formats []struct {
		URL string `json:"url"`
	} `json:"formats"`
}

func GetMusic(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, err := s.State.Guild(i.GuildID)
	if err != nil {
		fmt.Println(err)
	}

	if len(g.VoiceStates) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Music",
						Description: "Please connect to a voice channel to play music",
					},
				},
			},
		})
		return
	}

	url := i.ApplicationCommandData().Options[0].StringValue()
	cmd := exec.Command("youtube-dl", url, "--skip-download", "--print-json")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	yt := Youtube{}
	err = json.Unmarshal(stdout, &yt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go func() {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Music",
						Description: "Playing " + yt.Title,
					},
				},
			},
		})
	}()

	dgv, err := s.ChannelVoiceJoin(i.GuildID, g.VoiceStates[0].ChannelID, false, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	dgvoice.PlayAudioFile(dgv, yt.Formats[0].URL, make(<-chan bool))
}

func StopMusic(s *discordgo.Session, i *discordgo.InteractionCreate) {
	voice := s.VoiceConnections
	if len(voice) > 0 {
		go func() {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "Music",
							Description: "Stopped Playing Music!",
						},
					},
				},
			})
		}()

		voice[i.GuildID].Speaking(false)
		voice[i.GuildID].Disconnect()
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Music",
					Description: "Voice not connected!",
				},
			},
		},
	})
}

func PauseMusic(s *discordgo.Session, i *discordgo.InteractionCreate) {
	voice := s.VoiceConnections
	if len(voice) > 0 {
		voice[i.GuildID].Speaking(false)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Music",
						Description: "Music Paused!",
					},
				},
			},
		})

		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Music",
					Description: "Voice not connected!",
				},
			},
		},
	})
}

func FetchMusic(s *discordgo.Session, i *discordgo.InteractionCreate) {
	yt := ytdl.Client{
		Debug: false,
	}

	video, err := yt.GetVideo("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	if err != nil {
		fmt.Println(err)
	}
	audio := video.Formats.WithAudioChannels()
	fmt.Println(audio[0].Cipher)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Music",
					Description: "Music",
				},
			},
		},
	})

}
