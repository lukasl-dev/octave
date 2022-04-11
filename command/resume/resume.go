package resume

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/octave/command"
	"github.com/lukasl-dev/waterlink"
)

// Deps are the needed dependencies for Resume().
type Deps struct {
	// Conn is the waterlink connection to use for resuming the playback.
	Conn *waterlink.Connection `json:"conn,omitempty"`
}

// Resume returns the command to resumes the playback.
func Resume(deps Deps) command.Command {
	return command.Command{
		ApplicationCommand: discordgo.ApplicationCommand{
			Name:        "resume",
			Description: "Resumes the current playback",
		},
		Do: func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
			g := deps.Conn.Guild(snowflake.MustParse(evt.GuildID))

			if err := g.SetPaused(false); err != nil {
				return command.ErrorResponse(errors.New("failed to resume playback"))
			}

			return &discordgo.InteractionResponseData{
				Content: ":play_pause:  Playback has been resumed.",
			}
		},
	}
}
