package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rafmst/senhor-bot/db"
	"go.mongodb.org/mongo-driver/bson"
)

// HandleUnregister handles "!unregister" command
func HandleUnregister(session *discordgo.Session, command *discordgo.MessageCreate) {

	var user User
	users := db.Instance.Collection("users")

	filter := bson.M{"discord_id": command.Author.ID}
	err := users.FindOne(db.Ctx(), filter).Decode(&user)

	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "A sua informacão já foi previamente apagada")
	} else {
		_, err = users.DeleteOne(db.Ctx(), filter)
		if err == nil {
			session.ChannelMessageSend(command.ChannelID, "Toda a sua informacão foi apagada")
		}
	}
}
