package pause

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/octave/command"
	"github.com/lukasl-dev/waterlink"
)

// Deps are the needed dependencies for Pause().
type Deps struct {
	// Conn is the waterlink connection to use for pausing the playback.
	Conn *waterlink.Connection `json:"conn,omitempty"`
}

// Pause returns the command to pause the playback.
func Pause(deps Deps) command.Command {
	return command.Command{
		ApplicationCommand: discordgo.ApplicationCommand{
			Name:        "pause",
			Description: "Pauses the current playback.",
		},
		Do: func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
			g := deps.Conn.Guild(snowflake.MustParse(evt.GuildID))

			if err := g.SetPaused(true); err != nil {
				return command.ErrorResponse(errors.New("failed to pause playback"))
			}

			return &discordgo.InteractionResponseData{
				Content: ":pause_button: Playback has been paused.",
			}
		},
	}
}
