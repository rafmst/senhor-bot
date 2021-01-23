package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// Fox is the type of response received from FoxServer
type Fox struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}

// HandleFox handles "!fox" command
func HandleFox(session *discordgo.Session, command *discordgo.MessageCreate) {
	resp, err := http.Get("https://randomfox.ca/floof/")

	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "Error: FoxServer seems to be down")
	} else {

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Error: Wasn't able to read response")
		} else {
			var fox Fox
			err := json.Unmarshal(body, &fox)
			if err != nil {
				session.ChannelMessageSend(command.ChannelID, "Error: Check FoxServer response format")
			} else {
				session.ChannelMessageSend(command.ChannelID, fox.Image)
			}
		}
	}
}
