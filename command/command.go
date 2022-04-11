package command

import (
	"github.com/bwmarrin/discordgo"
)

// Command represents a (application) slash-command.
type Command struct {
	discordgo.ApplicationCommand

	// Do is the function to be called when the command is runned.
	Do func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData
}
