package seek

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lukasl-dev/octave/command"
	"github.com/lukasl-dev/waterlink/v2"
)

// Deps are the needed dependencies for Seek().
type Deps struct {
	// Conn is the waterlink connection to use for seeking the playback.
	Conn *waterlink.Connection
}

// Seek returns a command to seek the playback.
func Seek(deps Deps) command.Command {
	return command.Command{
		ApplicationCommand: discordgo.ApplicationCommand{
			Name:        "seek",
			Description: "Seeks to a specific position in the current playback.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "position",
					Description:  "The position to seek to.",
					Required:     true,
					Autocomplete: true,
				},
			},
		},
		Autocomplete: autocomplete(),
		Command:      run(deps),
	}
}
