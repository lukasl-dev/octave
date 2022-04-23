package guild

import "github.com/gompus/snowflake"

// Guild holds application related information about a guild.
type Guild struct {
	// ID is the unique identifier of the guild.
	ID snowflake.Snowflake `json:"id"`

	// Queue is the track.Track queue of the guild.
	Queue *Queue `json:"queue,omitempty"`
}
