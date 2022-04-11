package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gompus/snowflake"
	"log"
)

// registerHandlers registers all needed handlers for the current session.
func (a *app) registerHandlers() {
	a.session.AddHandler(a.handleReady)
	a.session.AddHandler(a.handleGuildJoin)
	a.session.AddHandler(a.handleVoiceServerUpdate)
	a.session.AddHandler(a.cmds.handle)
}

// handleReady handles incoming discordgo.Ready events. It creates all necessary
// components to interact with the lavalink server and registers the slash-commands
// on all guilds.
func (a *app) handleReady(s *discordgo.Session, evt *discordgo.Ready) {
	a.sessionID = evt.SessionID

	if err := a.createClient(); err != nil {
		log.Fatalln("Failed to create waterlink client:", err)
	}

	if err := a.createConnection(); err != nil {
		log.Fatalln("Failed to create waterlink connection:", err)
	}

	a.registerCommands()
	for _, guild := range a.session.State.Guilds {
		if err := a.cmds.create(a.session, guild.ID); err != nil {
			fmt.Printf("Failed to create slash-commands on guild %q: %s\n", guild.ID, err)
		}
	}

	log.Printf("Bot %s#%s is ready to serve %d guilds!\n", evt.User.Username, evt.User.Discriminator, len(evt.Guilds))
}

// handleGuildJoin handles incoming discordgo.GuildCreate events. It registers
// the slash-commands on the new guild.
func (a *app) handleGuildJoin(s *discordgo.Session, evt *discordgo.GuildCreate) {
	if err := a.cmds.create(s, evt.ID); err != nil {
		log.Printf("Failed to create slash-commands on guild %q: %s\n", evt.ID, err)
	}
}

// handleVoiceServerUpdate handles incoming discordgo.VoiceServerUpdate events.
func (a *app) handleVoiceServerUpdate(_ *discordgo.Session, evt *discordgo.VoiceServerUpdate) {
	g := a.conn.Guild(snowflake.MustParse(evt.GuildID))
	err := g.UpdateVoice(a.sessionID, evt.Token, evt.Endpoint)
	if err != nil {
		log.Printf("Failed to update voice server for guild %q: %s\n", evt.GuildID, err)
	}
}
