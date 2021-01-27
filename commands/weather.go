package commands

import (
	"os"
	"strconv"
	"strings"

	owm "github.com/briandowns/openweathermap"
	"github.com/bwmarrin/discordgo"
	"github.com/rafmst/senhor-bot/db"
	"go.mongodb.org/mongo-driver/bson"
)

// HandleWeather handles "!weather" command
func HandleWeather(session *discordgo.Session, command *discordgo.MessageCreate) {
	var (
		prefix     = os.Getenv("PREFIX")
		location   string
		usedPrefix string
	)

	if strings.HasPrefix(command.Content, prefix+"weather") {
		location = strings.TrimPrefix(command.Content, prefix+"weather")
		usedPrefix = prefix + "weather"
	} else {
		location = strings.TrimPrefix(command.Content, prefix+"w")
		usedPrefix = prefix + "w"
	}

	location = strings.TrimPrefix(location, " ")

	if len(location) > 0 {
		getWeatherInfo(location, session, command)
	} else {
		var user User
		users := db.Instance.Collection("users")

		filter := bson.M{"discord_id": command.Author.ID}
		err := users.FindOne(db.Ctx(), filter).Decode(&user)

		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Escolha uma localizacão, exemplo: `"+usedPrefix+" Mafamude`")
		} else {
			getWeatherInfo(user.City, session, command)
		}
	}
}

func getWeatherInfo(location string, session *discordgo.Session, command *discordgo.MessageCreate) {
	// Connect to weather api
	apiKey := os.Getenv("WEATHER_API_KEY")
	weather, err := owm.NewCurrent("C", "pt", apiKey)

	// Check for connection error
	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "Error: Connection to Weather API failed")
	}

	weather.CurrentByName(location)

	thumbnail := discordgo.MessageEmbedThumbnail{
		URL: "http://openweathermap.org/img/wn/" + weather.Weather[0].Icon + "@2x.png",
	}

	current := strconv.Itoa(int(weather.Main.Temp))
	max := strconv.Itoa(int(weather.Main.TempMax))
	min := strconv.Itoa(int(weather.Main.TempMin))

	var message = discordgo.MessageEmbed{
		Title: weather.Name,
		Description: `Temperatura actual: **` + current + `°C** 
Máximas de: **` + max + `°C**
Mínimas de: **` + min + `°C**`,
		Thumbnail: &thumbnail,
	}

	session.ChannelMessageSendEmbed(command.ChannelID, &message)
}
