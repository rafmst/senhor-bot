package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rafmst/senhor-bot/commands"
)

type commandHandler func(session *discordgo.Session, command *discordgo.MessageCreate)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func handleCommands(session *discordgo.Session, command *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if command.Author.ID == session.State.User.ID {
		return
	}

	var prefix = "!"

	if strings.HasPrefix(command.Content, prefix) {
		switch {
		case strings.HasPrefix(command.Content, prefix+"cat"):
			commands.HandleCat(session, command)
		}
	}
}
