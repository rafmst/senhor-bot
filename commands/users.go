package commands

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// HandleUsers handles "!users" command
func HandleUsers(session *discordgo.Session, command *discordgo.MessageCreate) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE")))
	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "Error: Couldn't connect to Mongo Database")
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Error: Couldn't establish connection to Mongo Database")
		} else {
			var usersFormatted string
			db := client.Database("senhor-bot")
			usersCollection := db.Collection("users")
			cursor, err := usersCollection.Find(ctx, bson.M{})
			if err != nil {
				session.ChannelMessageSend(command.ChannelID, "Error: No users found")
			} else {
				defer cursor.Close(ctx)
				for cursor.Next(ctx) {
					var user User
					if err = cursor.Decode(&user); err != nil {
						session.ChannelMessageSend(command.ChannelID, "Error: Reading user")
					} else {
						usersFormatted += `
│ ` + user.Name + fmt.Sprintf("%"+strconv.Itoa(11-len(user.Name))+"v", "") + `    │ ` + user.City + fmt.Sprintf("%"+strconv.Itoa(26-len(user.City))+"v", "") + ` │`
					}
				}

				session.ChannelMessageSend(command.ChannelID, `
`+"```"+`
┌────────────────┬────────────────────────────┐
│      User      │            City            │
├────────────────┼────────────────────────────┤ `+usersFormatted+`
└────────────────┴────────────────────────────┘
`+"```")
			}
		}
		cancel()
	}
}