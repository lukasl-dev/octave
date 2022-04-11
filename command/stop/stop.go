package stop

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/octave/command"
	"github.com/lukasl-dev/waterlink"
)

// Deps are the needed dependencies for Stop().
type Deps struct {
	// Conn is the waterlink connection to use for stopping the player.
	Conn *waterlink.Connection `json:"conn,omitempty"`
}

// Stop returns the command to stop a player.
func Stop(deps Deps) command.Command {
	return command.Command{
		ApplicationCommand: discordgo.ApplicationCommand{
			Name:        "stop",
			Description: "Stops the player",
		},
		Do: func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
			g := deps.Conn.Guild(snowflake.MustParse(evt.GuildID))

			if err := g.Stop(); err != nil {
				return command.ErrorResponse(errors.New("failed to stop the player"))
			}

			return &discordgo.InteractionResponseData{
				Content: ":stop_button: Player has been stopped.",
			}
		},
	}
}
