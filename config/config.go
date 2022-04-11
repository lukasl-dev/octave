package config

// Config contains configuration properties for the application.
type Config struct {
	// Token is the bot token to use.
	Token string `json:"token,omitempty"`

	// Lavalink is the lavalink configuration.
	Lavalink Lavalink `json:"lavalink,omitempty"`
}
