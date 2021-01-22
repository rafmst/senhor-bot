package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// Cat is the type of response received from CatServer
type Cat struct {
	File string `json:"file"`
}

// HandleCat handles "!cat" command
func HandleCat(session *discordgo.Session, command *discordgo.MessageCreate) {
	resp, err := http.Get("https://aws.random.cat/meow")

	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "Error: CatServer seems to be down")
	} else {

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Error: Wasn't able to read response")
		} else {
			var cat Cat
			err := json.Unmarshal(body, &cat)
			if err != nil {
				session.ChannelMessageSend(command.ChannelID, "Error: Check CatServer response format")
			} else {
				session.ChannelMessageSend(command.ChannelID, cat.File)
			}
		}
	}
}
