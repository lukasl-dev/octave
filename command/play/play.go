package play

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lukasl-dev/octave/command"
	"github.com/lukasl-dev/waterlink/v2"
)

// Deps are the needed dependencies for Play().
type Deps struct {
	// Client is the waterlink client to use to load tracks from a query.
	Client *waterlink.Client

	// Conn is the waterlink connection to use for playing music.
	Conn *waterlink.Connection
}

// Play returns a command to play music.
func Play(deps Deps) command.Command {
	return command.Command{
		ApplicationCommand: discordgo.ApplicationCommand{
			Name:        "play",
			Description: "Plays a song.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "query",
					Description:  "The query that is used to search for a song.",
					Required:     true,
					Autocomplete: true,
				},
			},
		},
		Autocomplete: autocomplete(),
		Command:      run(deps),
	}
}
