package commands

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

// HandleMowgli handles "!mowgli"
func HandleMowgli(session *discordgo.Session, command *discordgo.MessageCreate) {
	images := []string{
		"https://user-images.githubusercontent.com/924985/106322499-33397a80-6276-11eb-8e8b-d2e18c26538d.jpg",
		"https://user-images.githubusercontent.com/924985/106322504-346aa780-6276-11eb-8e2a-872a61b352a9.jpg",
		"https://user-images.githubusercontent.com/924985/106322508-346aa780-6276-11eb-9761-3c8529998ce8.jpg",
		"https://user-images.githubusercontent.com/924985/106322510-359bd480-6276-11eb-8fb5-abdb4dbbbbac.jpg",
		"https://user-images.githubusercontent.com/924985/106322513-36cd0180-6276-11eb-8164-6a71df0534cc.jpg",
		"https://user-images.githubusercontent.com/924985/106322514-37fe2e80-6276-11eb-9c82-8a0e2f35560e.jpg",
		"https://user-images.githubusercontent.com/924985/106322516-3896c500-6276-11eb-8bed-85bf88859130.jpg",
		"https://user-images.githubusercontent.com/924985/106322517-39c7f200-6276-11eb-833d-898df0d199f5.jpg",
		"https://user-images.githubusercontent.com/924985/106322519-3af91f00-6276-11eb-8037-3fb0d82283bb.jpg",
	}

	session.ChannelMessageSend(command.ChannelID, images[rand.Intn(len(images)-1)])
}
