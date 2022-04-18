package seek

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/octave/command"
	"github.com/lukasl-dev/waterlink/v2"
	"time"
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
		Command: func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
			position := command.Option(evt.ApplicationCommandData().Options, "position")

			parsed, err := time.ParseDuration(position.StringValue())
			if err != nil {
				return command.ErrorResponse(errors.New("invalid position"))
			}

			err = deps.Conn.Guild(snowflake.MustParse(evt.GuildID)).Seek(parsed)
			if err != nil {
				return command.ErrorResponse(errors.New("failed to seek"))
			}

			return &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(":hash: Seeked to **%s**.", position.StringValue()),
			}
		},
	}
}
