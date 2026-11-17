package scoring

import (
	"sync"
	"time"
)

// Presence says whether a user is in any voice channel.
type Presence int

const (
	OutOfVoice Presence = iota
	InVoice
)

// Stream says whether a user is sharing their screen (Go Live).
type Stream int

const (
	NotStreaming Stream = iota
	Streaming
)

// Earned is the whole minutes a closed voice or stream session is worth.
// Partial minutes are dropped.
type Earned struct {
	Voice     int
	Streaming int
}

type voiceSession struct {
	joined      time.Time
	streamStart time.Time
}

// VoiceTracker remembers when each user joined voice and started streaming.
// Sessions live in memory, so a restart drops the ones in progress.
type VoiceTracker struct {
	mu       sync.Mutex
	sessions map[string]*voiceSession
}

// NewVoiceTracker returns an empty tracker.
func NewVoiceTracker() *VoiceTracker {
	return &VoiceTracker{sessions: make(map[string]*voiceSession)}
}

func minutes(from, to time.Time) int {
	return int(to.Sub(from) / time.Minute)
}

// Update records a user's new voice state and returns what any session
// closed by this change earned. Moving between channels keeps the session.
func (v *VoiceTracker) Update(userID string, p Presence, s Stream, now time.Time) Earned {
	v.mu.Lock()
	defer v.mu.Unlock()

	sess, tracked := v.sessions[userID]
	if !tracked {
		if p == OutOfVoice {
			return Earned{}
		}
		sess = &voiceSession{joined: now}
		if s == Streaming {
			sess.streamStart = now
		}
		v.sessions[userID] = sess
		return Earned{}
	}

	var earned Earned
	streaming := !sess.streamStart.IsZero()

	if streaming && (p == OutOfVoice || s == NotStreaming) {
		earned.Streaming = minutes(sess.streamStart, now)
		sess.streamStart = time.Time{}
	}
	if !streaming && s == Streaming {
		sess.streamStart = now
	}

	if p == OutOfVoice {
		earned.Voice = minutes(sess.joined, now)
		delete(v.sessions, userID)
	}

	return earned
}
