package scoring

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

// InviteCache keeps each guild's invite use counts so a new member can be
// traced to the invite they used. Discord does not say which one it was.
type InviteCache struct {
	mu     sync.Mutex
	guilds map[string]map[string]int
}

// NewInviteCache returns an empty cache.
func NewInviteCache() *InviteCache {
	return &InviteCache{guilds: make(map[string]map[string]int)}
}

func useCounts(invites []*discordgo.Invite) map[string]int {
	uses := make(map[string]int, len(invites))
	for _, inv := range invites {
		uses[inv.Code] = inv.Uses
	}
	return uses
}

// Set replaces a guild's cached invites.
func (c *InviteCache) Set(guildID string, invites []*discordgo.Invite) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.guilds[guildID] = useCounts(invites)
}

// Used compares fresh invites against the cache, stores the fresh counts,
// and returns the inviter whose invite gained a use. An invite missing from
// the cache counts as used once it has any uses. Returns nil when no invite
// changed, the guild had no baseline yet, or a single use invite was
// already deleted by Discord.
func (c *InviteCache) Used(guildID string, fresh []*discordgo.Invite) *discordgo.User {
	c.mu.Lock()
	defer c.mu.Unlock()

	cached, known := c.guilds[guildID]
	c.guilds[guildID] = useCounts(fresh)
	if !known {
		return nil
	}

	for _, inv := range fresh {
		if inv.Inviter == nil || inv.Uses <= cached[inv.Code] {
			continue
		}
		return inv.Inviter
	}

	return nil
}

// Invites is the bot wide invite cache.
var Invites = NewInviteCache()
