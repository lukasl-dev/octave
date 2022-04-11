package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lukasl-dev/octave/command"
)

// commandHandler is responsible for handling slash-commands.
type commandHandler struct {
	// commands is a map of all the commands that can be called using
	// slash-commands.
	commands map[string]command.Command

	appCmds []*discordgo.ApplicationCommand
}

// newCommandHandler returns a new command handler that is responsible for
// routing slash-commands to the given commands.
func newCommandHandler() *commandHandler {
	return &commandHandler{commands: make(map[string]command.Command)}
}

// add registers the given command in the handler.
func (c *commandHandler) add(cmd command.Command) {
	c.commands[cmd.Name] = cmd
	c.appCmds = append(c.appCmds, &cmd.ApplicationCommand)
}

// create creates all registered commands on the guild whose id is given.
func (c *commandHandler) create(s *discordgo.Session, gID string) error {
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, gID, c.appCmds)
	return err
}

// handle processes incoming interaction create events and calls the appropriate
// command.
func (c *commandHandler) handle(s *discordgo.Session, evt *discordgo.InteractionCreate) {
	cmd := c.commands[evt.ApplicationCommandData().Name]
	_ = s.InteractionRespond(evt.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: cmd.Do(s, evt),
	})
}
