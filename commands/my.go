package commands

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// HandleMyCity handles "!mycity" command
func HandleMyCity(session *discordgo.Session, command *discordgo.MessageCreate) {
	var (
		prefix = os.Getenv("PREFIX")
		city   string
	)

	city = strings.TrimPrefix(command.Content, prefix+"mycity")
	city = strings.TrimPrefix(city, " ")

	if len(city) > 0 {
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
					session.ChannelMessageSend(command.ChannelID, "Utilizador não existente, registe-se com `!register`")
				} else {
					_, err := users.UpdateOne(ctx, filter, bson.D{
						{Key: "$set", Value: bson.M{"city": city}},
					})
					if err != nil {
						session.ChannelMessageSend(command.ChannelID, "Erro ao alterar a sua cidade default")
					} else {
						session.ChannelMessageSend(command.ChannelID, "Cidade alterada com sucesso")
					}
				}
			}
			cancel()
		}
	} else {
		session.ChannelMessageSend(command.ChannelID, "Escolha uma localizacão, exemplo: `!mycity Mafamude`")
	}
}
