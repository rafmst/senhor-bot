package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rafmst/senhor-bot/db"
	"go.mongodb.org/mongo-driver/bson"
)

// User type from database
type User struct {
	DiscordID string `bson:"discord_id"`
	Name      string `bson:"name"`
	City      string `bson:"city"`
}

// HandleRegister handles "!register" command
func HandleRegister(session *discordgo.Session, command *discordgo.MessageCreate) {
	var user User

	users := db.Instance.Collection("users")
	filter := bson.M{"discord_id": command.Author.ID}
	err := users.FindOne(db.Ctx(), filter).Decode(&user)

	if err != nil {
		newUser := User{command.Author.ID, command.Author.Username, ""}
		_, err := users.InsertOne(db.Ctx(), newUser)
		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Erro ao inserir utilizador")
		} else {
			session.ChannelMessageSend(command.ChannelID, "Utilizador registado com sucesso")
		}
	} else {
		session.ChannelMessageSend(command.ChannelID, "Este utilizador já está registado")
	}
}
