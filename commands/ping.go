package commands

import "github.com/bwmarrin/discordgo"

// HandlePing handles "!ping" command
func HandlePing(session *discordgo.Session, command *discordgo.MessageCreate) {
	session.ChannelMessageSend(command.ChannelID, "Pong!")
}
