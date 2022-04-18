package command

import (
	"github.com/bwmarrin/discordgo"
)

// Command represents a (application) slash-command.
type Command struct {
	discordgo.ApplicationCommand

	// Autocomplete is the Handler to be called when the user requests auto-completion.
	Autocomplete Handler

	// Command is the Handler to be called when the command is run.
	Command Handler
}

// Handler is a function that handles a discordgo.InteractionCreate event.
type Handler = func(s *discordgo.Session, evt *discordgo.InteractionCreate) *discordgo.InteractionResponseData
