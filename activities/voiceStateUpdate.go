package activities

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dreygur/leaderboardbot/hooks"
	"github.com/dreygur/leaderboardbot/lib"
	"github.com/dreygur/leaderboardbot/repo"
	"github.com/dreygur/leaderboardbot/scoring"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var voiceSessions = scoring.NewVoiceTracker()

func voiceState(v *discordgo.VoiceState) (scoring.Presence, scoring.Stream) {
	presence, stream := scoring.OutOfVoice, scoring.NotStreaming
	if v.ChannelID != "" {
		presence = scoring.InVoice
	}
	if v.SelfStream {
		stream = scoring.Streaming
	}
	return presence, stream
}

func awardVoice(s *discordgo.Session, userID string, earned scoring.Earned) {
	if earned == (scoring.Earned{}) {
		return
	}

	update := bson.M{
		"points":           lib.POINTS["voice"]*earned.Voice + lib.POINTS["streaming"]*earned.Streaming,
		"activities.voice": earned.Voice,
	}

	level, err := repo.Award(userID, update)
	if err != nil {
		lib.PrintLog(fmt.Sprintf("Error awarding voice points: %v", err), "error")
	}
	if level > 0 {
		hooks.DMLevel(s, userID, level)
	}
}

func VoiceStateUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	if v.UserID == s.State.User.ID {
		return
	}

	presence, stream := voiceState(v.VoiceState)
	awardVoice(s, v.UserID, voiceSessions.Update(v.UserID, presence, stream, time.Now()))

	user, err := s.State.Member(v.GuildID, v.UserID)
	if err != nil {
		lib.PrintLog("Failed to get user name for "+v.UserID, "error")
		return
	}
	lib.PrintLog(fmt.Sprintf("VoiceStateUpdate from %s: %s", user.User.Username, v.ChannelID), "info")
}
