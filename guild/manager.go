package guild

import (
	"github.com/gompus/snowflake"
	"sync"
)

// Manager manages all guilds on which the bot is present.
type Manager struct {
	// mu protects the following fields from concurrent access.
	mu sync.RWMutex

	// guilds is a map of guilds on which the bot is operating on.
	guilds map[snowflake.Snowflake]*Guild
}

// NewManager returns a new Manager with an empty guilds map.
func NewManager() *Manager {
	return &Manager{guilds: make(map[snowflake.Snowflake]*Guild)}
}

// Range calls fn() sequentially for each guild until fn() returns false.
func (m *Manager) Range(fn func(g *Guild) bool) {
	m.mu.Lock()
	for _, guild := range m.guilds {
		if !fn(guild) {
			break
		}
	}
	m.mu.Unlock()
}

// Guilds returns a slice of all guilds on which the bot is present.
func (m *Manager) Guilds() []*Guild {
	guilds := make([]*Guild, 0, len(m.guilds))

	m.mu.RLock()
	for _, guild := range m.guilds {
		guilds = append(guilds, guild)
	}
	m.mu.RUnlock()

	return guilds
}

// Guild searches for a guild by its ID and returns it, or nil if not found.
func (m *Manager) Guild(id snowflake.Snowflake) *Guild {
	m.mu.RLock()
	guild, ok := m.guilds[id]
	m.mu.RUnlock()

	if !ok {
		return nil
	}
	return guild
}

// Add adds a guild to the manager.
func (m *Manager) Add(guild *Guild) {
	m.mu.Lock()
	m.guilds[guild.ID] = guild
	m.mu.Unlock()
}

// Delete removes a guild from the manager.
func (m *Manager) Delete(id snowflake.Snowflake) {
	m.mu.Lock()
	delete(m.guilds, id)
	m.mu.Unlock()
}
