package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// Dog is the type of response received from DogServer
type Dog struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// HandleDog handles "!cat" command
func HandleDog(session *discordgo.Session, command *discordgo.MessageCreate) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")

	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "Error: DogServer seems to be down")
	} else {

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Error: Wasn't able to read response")
		} else {
			var dog Dog
			err := json.Unmarshal(body, &dog)
			if err != nil {
				session.ChannelMessageSend(command.ChannelID, "Error: Check DogServer response format")
			} else {
				session.ChannelMessageSend(command.ChannelID, dog.Message)
			}
		}
	}
}
