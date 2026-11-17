package scoring

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func invite(code, inviterID string, uses int) *discordgo.Invite {
	return &discordgo.Invite{
		Code:    code,
		Uses:    uses,
		Inviter: &discordgo.User{ID: inviterID},
	}
}

func TestInviteUsedFindsBumpedCode(t *testing.T) {
	c := NewInviteCache()
	c.Set("g1", []*discordgo.Invite{invite("aaa", "alice", 2), invite("bbb", "bob", 5)})

	got := c.Used("g1", []*discordgo.Invite{invite("aaa", "alice", 2), invite("bbb", "bob", 6)})

	if got == nil || got.ID != "bob" {
		t.Fatalf("got %v, want bob", got)
	}
}

func TestInviteUsedSeesInviteCreatedAfterCache(t *testing.T) {
	c := NewInviteCache()
	c.Set("g1", []*discordgo.Invite{invite("aaa", "alice", 2)})

	got := c.Used("g1", []*discordgo.Invite{invite("aaa", "alice", 2), invite("new", "carol", 1)})

	if got == nil || got.ID != "carol" {
		t.Fatalf("got %v, want carol", got)
	}
}

func TestInviteUsedRefreshesCache(t *testing.T) {
	c := NewInviteCache()
	c.Set("g1", []*discordgo.Invite{invite("aaa", "alice", 2)})

	c.Used("g1", []*discordgo.Invite{invite("aaa", "alice", 3)})
	again := c.Used("g1", []*discordgo.Invite{invite("aaa", "alice", 3)})

	if again != nil {
		t.Fatalf("got %v, want nil once cache holds the new count", again)
	}
}

func TestInviteUsedNoChange(t *testing.T) {
	c := NewInviteCache()
	c.Set("g1", []*discordgo.Invite{invite("aaa", "alice", 2)})

	if got := c.Used("g1", []*discordgo.Invite{invite("aaa", "alice", 2)}); got != nil {
		t.Fatalf("got %v, want nil", got)
	}
}

func TestInviteUsedSkipsMissingInviter(t *testing.T) {
	c := NewInviteCache()
	c.Set("g1", nil)
	vanity := &discordgo.Invite{Code: "vanity", Uses: 1}

	if got := c.Used("g1", []*discordgo.Invite{vanity}); got != nil {
		t.Fatalf("got %v, want nil for invite with no inviter", got)
	}
}

func TestInviteUsedUncachedGuildGuessesNothing(t *testing.T) {
	c := NewInviteCache()

	first := c.Used("g1", []*discordgo.Invite{invite("aaa", "alice", 4)})
	second := c.Used("g1", []*discordgo.Invite{invite("aaa", "alice", 5)})

	if first != nil {
		t.Fatalf("first join got %v, want nil with no baseline", first)
	}
	if second == nil || second.ID != "alice" {
		t.Fatalf("second join got %v, want alice", second)
	}
}
