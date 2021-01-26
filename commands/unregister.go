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

// HandleUnregister handles "!unregister" command
func HandleUnregister(session *discordgo.Session, command *discordgo.MessageCreate) {
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
				session.ChannelMessageSend(command.ChannelID, "A sua informacão já foi previamente apagada")
			} else {
				_, err = users.DeleteOne(ctx, filter)
				if err == nil {
					session.ChannelMessageSend(command.ChannelID, "Toda a sua informacão foi apagada")
				}
			}
		}
		cancel()
	}
}
