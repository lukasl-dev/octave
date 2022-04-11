package volume

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/octave/command"
	"github.com/lukasl-dev/waterlink"
)

// Deps are the needed dependencies for Volume().
type Deps struct {
	// Conn is the waterlink connection to use for updating the volume.
	Conn *waterlink.Connection `json:"conn,omitempty"`
}

// Volume returns the command to update the volume.
func Volume(deps Deps) command.Command {
	return command.Command{
		ApplicationCommand: discordgo.ApplicationCommand{
			Name:        "volume",
			Description: "Update the player's volume",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "volume",
					Description:  "The new volume in percent",
					MaxValue:     100,
					Type:         discordgo.ApplicationCommandOptionInteger,
					Autocomplete: true,
					Required:     true,
				},
			},
		},
		Do: func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
			volume := command.Option(evt.ApplicationCommandData().Options, "volume")

			g := deps.Conn.Guild(snowflake.MustParse(evt.GuildID))
			if err := g.UpdateVolume(uint16(volume.IntValue())); err != nil {
				return command.ErrorResponse(errors.New("failed to update volume"))
			}

			return &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(":sound: Volume has been updated to %d%%.", volume.IntValue()),
			}
		},
	}
}
