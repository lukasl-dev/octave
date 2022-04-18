package play

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/lukasl-dev/octave/command"
	"github.com/lukasl-dev/waterlink/v2/track/query"
	"net/url"
	"strings"
)

func autocomplete() command.Handler {
	return func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
		q := command.Option(evt.ApplicationCommandData().Options, "query")

		var choices []*discordgo.ApplicationCommandOptionChoice
		if len(strings.TrimSpace(q.StringValue())) > 0 {
			_, err := url.ParseRequestURI(q.StringValue())
			if err == nil {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  fmt.Sprintf("\U0001F7E6 URL: %s", q.StringValue()),
					Value: q.StringValue(),
				})
			}

			choices = append(choices,
				&discordgo.ApplicationCommandOptionChoice{
					Name:  fmt.Sprintf("\U0001F7E5 YouTube: %s", q.StringValue()),
					Value: query.YouTube(q.StringValue()),
				},
				&discordgo.ApplicationCommandOptionChoice{
					Name:  fmt.Sprintf("\U0001F7E7 SoundCloud: %s", q.StringValue()),
					Value: query.SoundCloud(q.StringValue()),
				},
			)
		}

		return &discordgo.InteractionResponseData{Choices: choices}
	}
}
