package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Dog is the type of response received from DogServer
type Dog struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Fox is the type of response received from FoxServer
type Fox struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}

// Cat is the type of response received from CatServer
type Cat struct {
	File string `json:"file"`
}

// AnimalAPIUrls saves all the urls for the different random animals APIs
type AnimalAPIUrls struct {
	Cat string
	Dog string
	Fox string
}

var animalAPIUrls = AnimalAPIUrls{
	Cat: "https://aws.random.cat/meow",
	Dog: "https://dog.ceo/api/breeds/image/random",
	Fox: "https://randomfox.ca/floof/",
}

// HandleAnimal handles "!dog", "!fox" and "!cat" command
func HandleAnimal(session *discordgo.Session, command *discordgo.MessageCreate) {
	prefix := os.Getenv("PREFIX")
	content := command.Content
	var url string

	switch {
	case strings.HasPrefix(content, prefix+"dog"):
		url = animalAPIUrls.Dog
	case strings.HasPrefix(content, prefix+"fox"):
		url = animalAPIUrls.Fox
	case strings.HasPrefix(content, prefix+"cat"):
		url = animalAPIUrls.Cat
	}

	resp, err := http.Get(url)

	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "Error: Server seems to be down")
	} else {

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Error: Wasn't able to read response")
		} else {
			var image string
			switch {
			case strings.HasPrefix(content, prefix+"dog"):
				image, err = handleDog(body)
			case strings.HasPrefix(content, prefix+"fox"):
				image, err = handleFox(body)
			case strings.HasPrefix(content, prefix+"cat"):
				image, err = handleCat(body)
			}

			if err != nil {
				session.ChannelMessageSend(command.ChannelID, "Error: Check DogServer response format")
			} else {
				session.ChannelMessageSend(command.ChannelID, image)
			}
		}
	}
}

func handleDog(body []byte) (string, error) {
	var result string
	var dog Dog
	err := json.Unmarshal(body, &dog)
	if err != nil {
		return result, err
	} else {
		return dog.Message, nil
	}
}

func handleCat(body []byte) (string, error) {
	var result string
	var cat Cat
	err := json.Unmarshal(body, &cat)
	if err != nil {
		return result, err
	} else {
		return cat.File, nil
	}
}

func handleFox(body []byte) (string, error) {
	var result string
	var fox Fox
	err := json.Unmarshal(body, &fox)
	if err != nil {
		return result, err
	} else {
		return fox.Image, nil
	}
}
