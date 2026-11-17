package scoring

import (
	"testing"
	"time"
)

var start = time.Date(2026, 11, 17, 12, 0, 0, 0, time.UTC)

func at(d time.Duration) time.Time {
	return start.Add(d)
}

func TestVoiceLeaveEarnsWholeMinutes(t *testing.T) {
	v := NewVoiceTracker()

	v.Update("u1", InVoice, NotStreaming, at(0))
	got := v.Update("u1", OutOfVoice, NotStreaming, at(5*time.Minute+59*time.Second))

	if got != (Earned{Voice: 5}) {
		t.Fatalf("got %+v, want 5 voice minutes", got)
	}
}

func TestVoiceMoveKeepsSession(t *testing.T) {
	v := NewVoiceTracker()

	v.Update("u1", InVoice, NotStreaming, at(0))
	moved := v.Update("u1", InVoice, NotStreaming, at(3*time.Minute))
	left := v.Update("u1", OutOfVoice, NotStreaming, at(10*time.Minute))

	if moved != (Earned{}) {
		t.Fatalf("move earned %+v, want nothing", moved)
	}
	if left != (Earned{Voice: 10}) {
		t.Fatalf("leave earned %+v, want 10 voice minutes", left)
	}
}

func TestVoiceStreamStopEarnsStreaming(t *testing.T) {
	v := NewVoiceTracker()

	v.Update("u1", InVoice, NotStreaming, at(0))
	v.Update("u1", InVoice, Streaming, at(2*time.Minute))
	stopped := v.Update("u1", InVoice, NotStreaming, at(9*time.Minute))
	left := v.Update("u1", OutOfVoice, NotStreaming, at(12*time.Minute))

	if stopped != (Earned{Streaming: 7}) {
		t.Fatalf("stream stop earned %+v, want 7 streaming minutes", stopped)
	}
	if left != (Earned{Voice: 12}) {
		t.Fatalf("leave earned %+v, want 12 voice minutes", left)
	}
}

func TestVoiceLeaveWhileStreaming(t *testing.T) {
	v := NewVoiceTracker()

	v.Update("u1", InVoice, Streaming, at(0))
	got := v.Update("u1", OutOfVoice, NotStreaming, at(4*time.Minute))

	if got != (Earned{Voice: 4, Streaming: 4}) {
		t.Fatalf("got %+v, want 4 voice and 4 streaming minutes", got)
	}
}

func TestVoiceUnknownLeaveEarnsNothing(t *testing.T) {
	v := NewVoiceTracker()

	got := v.Update("u1", OutOfVoice, NotStreaming, at(time.Hour))

	if got != (Earned{}) {
		t.Fatalf("got %+v, want nothing", got)
	}
}

func TestVoiceSessionEndsOnLeave(t *testing.T) {
	v := NewVoiceTracker()

	v.Update("u1", InVoice, NotStreaming, at(0))
	v.Update("u1", OutOfVoice, NotStreaming, at(time.Minute))
	again := v.Update("u1", OutOfVoice, NotStreaming, at(time.Hour))

	if again != (Earned{}) {
		t.Fatalf("second leave earned %+v, want nothing", again)
	}
}
