package commands

import (
	"context"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User type from database
type User struct {
	DiscordID string `bson:"discord_id"`
	Name      string `bson:"name"`
	City      string `bson:"city"`
}

// HandleRegister handles "!register" command
func HandleRegister(session *discordgo.Session, command *discordgo.MessageCreate) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE")))
	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "Error: Couldn't connect to Mongo Database")
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Error: Couldn't establish connection to Mongo Database")
		} else {
			var user User
			db := client.Database("senhor-bot")
			users := db.Collection("users")

			filter := bson.M{"discord_id": command.Author.ID}
			err = users.FindOne(ctx, filter).Decode(&user)

			if err != nil {
				newUser := User{command.Author.ID, command.Author.Username, ""}
				_, err := users.InsertOne(ctx, newUser)
				if err != nil {
					session.ChannelMessageSend(command.ChannelID, "Erro ao inserir utilizador")
				} else {
					session.ChannelMessageSend(command.ChannelID, "Utilizador registado com sucesso")
				}
			} else {
				session.ChannelMessageSend(command.ChannelID, "Este utilizador já está registado")
			}
		}
		cancel()
	}
}
