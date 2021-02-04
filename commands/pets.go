package commands

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// HandlePets handles "!mowgli", "!napoleao" & "!safira"
func HandlePets(session *discordgo.Session, command *discordgo.MessageCreate) {
	prefix := os.Getenv("PREFIX")
	content := command.Content

	pet := strings.TrimPrefix(content, prefix)
	pet = strings.Split(pet, " ")[0]

	images := map[string][]string{
		"mowgli": {
			"https://user-images.githubusercontent.com/924985/106322499-33397a80-6276-11eb-8e8b-d2e18c26538d.jpg",
			"https://user-images.githubusercontent.com/924985/106322504-346aa780-6276-11eb-8e2a-872a61b352a9.jpg",
			"https://user-images.githubusercontent.com/924985/106322508-346aa780-6276-11eb-9761-3c8529998ce8.jpg",
			"https://user-images.githubusercontent.com/924985/106322510-359bd480-6276-11eb-8fb5-abdb4dbbbbac.jpg",
			"https://user-images.githubusercontent.com/924985/106322513-36cd0180-6276-11eb-8164-6a71df0534cc.jpg",
			"https://user-images.githubusercontent.com/924985/106322514-37fe2e80-6276-11eb-9c82-8a0e2f35560e.jpg",
			"https://user-images.githubusercontent.com/924985/106322516-3896c500-6276-11eb-8bed-85bf88859130.jpg",
			"https://user-images.githubusercontent.com/924985/106322517-39c7f200-6276-11eb-833d-898df0d199f5.jpg",
			"https://user-images.githubusercontent.com/924985/106322519-3af91f00-6276-11eb-8037-3fb0d82283bb.jpg",
		},
		"napoleao": {
			"https://user-images.githubusercontent.com/924985/106795135-00b6c580-665a-11eb-8899-7e217d7334f2.png",
			"https://user-images.githubusercontent.com/924985/106795538-7d49a400-665a-11eb-8380-7c4fccec1374.png",
			"https://user-images.githubusercontent.com/924985/106795569-88043900-665a-11eb-8627-61ee7b6921ba.png",
			"https://user-images.githubusercontent.com/924985/106795605-918da100-665a-11eb-88e3-0fc1ffa9b754.png",
			"https://user-images.githubusercontent.com/924985/106795637-99e5dc00-665a-11eb-9275-e4ebb4afce95.png",
			"https://user-images.githubusercontent.com/924985/106796078-301a0200-665b-11eb-8aed-71ba4192906b.png",
			"https://user-images.githubusercontent.com/924985/106796102-360fe300-665b-11eb-93f9-6d67994c4296.png",
			"https://user-images.githubusercontent.com/924985/106796148-4922b300-665b-11eb-86db-8ce1ae904a73.png",
			"https://user-images.githubusercontent.com/924985/106796245-66f01800-665b-11eb-9535-2d69ded4f181.png",
		},
		"safira": {
			"https://user-images.githubusercontent.com/924985/106796542-c9e1af00-665b-11eb-9535-ca74101d088d.png",
			"https://user-images.githubusercontent.com/924985/106797569-124d9c80-665d-11eb-83c9-1c91a40113c3.png",
			"https://user-images.githubusercontent.com/924985/106925715-a2005300-6710-11eb-8b7d-696590454c16.png",
			"https://user-images.githubusercontent.com/924985/106925777-b2b0c900-6710-11eb-9100-b901203ebf0b.png",
			"https://user-images.githubusercontent.com/924985/106925890-cc521080-6710-11eb-951c-47075be616f8.png",
			"https://user-images.githubusercontent.com/924985/106925930-d6740f00-6710-11eb-89ed-313380bb9938.png",
			"https://user-images.githubusercontent.com/924985/106925962-e25fd100-6710-11eb-8f1f-f3f7f77d5aae.png",
			"https://user-images.githubusercontent.com/924985/106926003-ee4b9300-6710-11eb-8333-ec9e03e6a2bf.png",
			"https://user-images.githubusercontent.com/924985/106926072-fe637280-6710-11eb-9410-f9b54a2dde46.png",
		},
	}

	rand.Seed(time.Now().Unix())
	session.ChannelMessageSend(command.ChannelID, images[pet][rand.Intn(len(images[pet])-1)])
}
