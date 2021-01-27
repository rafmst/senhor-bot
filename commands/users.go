package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rafmst/senhor-bot/db"
	"go.mongodb.org/mongo-driver/bson"
)

// HandleUsers handles "!users" command
func HandleUsers(session *discordgo.Session, command *discordgo.MessageCreate) {
	usersCollection := db.Instance.Collection("users")
	cursor, err := usersCollection.Find(db.Ctx(), bson.M{})
	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "Error: No users found")
	} else {
		defer cursor.Close(db.Ctx())
		var usersFormatted string
		for cursor.Next(db.Ctx()) {
			var user User
			if err = cursor.Decode(&user); err != nil {
				session.ChannelMessageSend(command.ChannelID, "Error: Reading user")
			} else {
				usersFormatted += `
│` + calculateSpacesSufix(15, user.Name) + `│` + calculateSpacesSufix(19, user.City) + `│`
			}
		}

		session.ChannelMessageSend(command.ChannelID, `
`+"```"+`
┌────────────────┬────────────────────┐
│      User      │        City        │
├────────────────┼────────────────────┤ `+usersFormatted+`
└────────────────┴────────────────────┘
`+"```")
	}
}

func calculateSpacesSufix(width int, text string) string {
	var spaces string
	for missingSpaces := width - len(text); missingSpaces > 0; missingSpaces-- {
		spaces += " "
	}

	return " " + text + spaces
}
