package play

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/octave/command"
	"github.com/lukasl-dev/waterlink"
	"github.com/lukasl-dev/waterlink/track/query"
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
			Description: "Plays a song from a given URL",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "url",
					Description:  "The URL of the song to play",
					Autocomplete: true,
					Required:     true,
				},
			},
		},
		Do: func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
			url := command.Option(evt.ApplicationCommandData().Options, "url")

			res, err := deps.Client.LoadTracks(query.Of(url.StringValue()))
			switch {
			case err != nil:
				return command.ErrorResponse(errors.New("failed to load track"))
			case len(res.Tracks) == 0:
				return command.ErrorResponse(errors.New("no tracks found"))
			}
			tr := res.Tracks[0]

			channelID := command.MemberChannel(s, evt.GuildID, evt.Member.User.ID)
			if channelID == "" {
				return command.ErrorResponse(errors.New("user is not a member of any voice channel"))
			}

			err = s.ChannelVoiceJoinManual(evt.GuildID, channelID, false, true)
			if err != nil {
				return command.ErrorResponse(errors.New("failed to join voice channel"))
			}

			err = deps.Conn.Guild(snowflake.MustParse(evt.GuildID)).PlayTrack(tr)
			if err != nil {
				return command.ErrorResponse(errors.New("failed to play track"))
			}

			return &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(":arrow_forward: Now playing **%s**", tr.Info.Title),
			}
		},
	}
}
