package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gompus/snowflake"
	"github.com/lukasl-dev/octave/command/pause"
	"github.com/lukasl-dev/octave/command/play"
	"github.com/lukasl-dev/octave/config"
	"github.com/lukasl-dev/waterlink"
)

// app is the main application struct. It holds all necessary dependencies to
// run the bot.
type app struct {
	// cfg is the application configuration.
	cfg config.Config

	// session is the discord session on which the bot is running on.
	session *discordgo.Session

	// conn is the active lavalink connection.
	conn *waterlink.Connection

	// client is the lavalink client.
	client *waterlink.Client

	// cmds is the commandHandler that is responsible for handling commands.
	cmds *commandHandler

	// sessionID is the current session's ID which is received on discordgo.Ready.
	sessionID string
}

// newApp returns a new app configured by the given config.Config.
func newApp(cfg config.Config) *app {
	return &app{cfg: cfg, cmds: newCommandHandler()}
}

func (a *app) run() (err error) {
	a.session, err = discordgo.New(fmt.Sprintf("Bot %s", a.cfg.Token))
	if err != nil {
		return fmt.Errorf("failed to create discord session: %w", err)
	}

	a.registerHandlers()
	return a.session.Open()
}

// createClient tries to create a new waterlink.Client and defines it in the app.
func (a *app) createClient() (err error) {
	a.client, err = waterlink.NewClient(fmt.Sprintf("http://%s", a.cfg.Lavalink.Host), a.credentials())
	return err
}

// createConnection tries to create a new waterlink.Connection and defines it in
// the app.
func (a *app) createConnection() (err error) {
	a.conn, err = waterlink.Open(fmt.Sprintf("ws://%s", a.cfg.Lavalink.Host), a.credentials())
	return err
}

// credentials returns the waterlink.Credentials to use for client and connection.
func (a *app) credentials() waterlink.Credentials {
	return waterlink.Credentials{
		Authorization: a.cfg.Lavalink.Passphrase,
		UserID:        snowflake.MustParse(a.session.State.User.ID),
	}
}

// registerCommands registers all commands in the internal commandHandler.
func (a *app) registerCommands() {
	a.cmds.add(play.Play(play.Deps{Client: a.client, Conn: a.conn}))
	a.cmds.add(pause.Pause(pause.Deps{Conn: a.conn}))
}
