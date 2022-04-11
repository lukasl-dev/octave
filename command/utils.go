package command

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Option searches the given options by its name and returns it.
func Option(opts []*discordgo.ApplicationCommandInteractionDataOption, name string) *discordgo.ApplicationCommandInteractionDataOption {
	for _, opt := range opts {
		if strings.EqualFold(opt.Name, name) {
			return opt
		}
	}
	return nil
}

// MemberChannel returns the id of the current channel the user	is in or and
// empty string if the user is not in a channel.
func MemberChannel(s *discordgo.Session, guildID, memberID string) string {
	guild, err := s.State.Guild(guildID)
	if err != nil {
		return ""
	}

	for _, state := range guild.VoiceStates {
		if strings.EqualFold(memberID, state.UserID) {
			return state.ChannelID
		}
	}

	return ""
}
