package seek

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/octave/command"
	"time"
)

func run(deps Deps) command.Handler {
	return func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
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
	}
}
