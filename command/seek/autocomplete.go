package seek

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lukasl-dev/octave/command"
	"unicode"
)

func autocomplete() command.Handler {
	return func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
		position := command.Option(evt.ApplicationCommandData().Options, "position")

		var val string
		for _, r := range position.StringValue() {
			if unicode.IsDigit(r) {
				val += string(r)
			}
		}

		if len(val) == 0 {
			val = "0"
		}

		return &discordgo.InteractionResponseData{
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  val + "ms",
					Value: val + "ms",
				},
				{
					Name:  val + "s",
					Value: val + "s",
				},
				{
					Name:  val + "m",
					Value: val + "m",
				},
				{
					Name:  val + "h",
					Value: val + "h",
				},
			},
		}
	}
}
