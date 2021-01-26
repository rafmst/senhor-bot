package main

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rafmst/senhor-bot/commands"
)

type commandHandler func(session *discordgo.Session, command *discordgo.MessageCreate)

var content string

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func handleCommands(session *discordgo.Session, command *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if command.Author.ID == session.State.User.ID {
		return
	}

	content = command.Content
	if hasPrefix("") {
		switch {
		case hasPrefix("help"):
			commands.HandleHelp(session, command)
		case hasPrefix("cat"):
			commands.HandleCat(session, command)
		case hasPrefix("dog"):
			commands.HandleDog(session, command)
		case hasPrefix("fox"):
			commands.HandleFox(session, command)
		case hasPrefix("weather") || hasPrefix("w"):
			commands.HandleWeather(session, command)
		case hasPrefix("covid"):
			commands.HandleCovid(session, command)
		case hasPrefix("register"):
			commands.HandleRegister(session, command)
		case hasPrefix("mycity"):
			commands.HandleMyCity(session, command)
		case hasPrefix("unregister"):
			commands.HandleUnregister(session, command)
		case hasPrefix("users"):
			commands.HandleUsers(session, command)
		}
	}
}

func hasPrefix(keyword string) bool {
	prefix := os.Getenv("PREFIX")
	return strings.HasPrefix(content, prefix+keyword)
}
