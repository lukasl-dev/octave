package config

// Lavalink holds lavalink configuration values.
type Lavalink struct {
	// Host is the host of the lavalink server instance. For instance: 'localhost:2333'
	Host string `json:"host,omitempty"`

	// Passphrase is the passphrase to use for authentication with the lavalink
	// server.
	Passphrase string `json:"passphrase,omitempty"`
}
