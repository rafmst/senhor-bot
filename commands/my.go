package commands

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rafmst/senhor-bot/db"
	"go.mongodb.org/mongo-driver/bson"
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
		var user User
		users := db.Instance.Collection("users")

		filter := bson.M{"discord_id": command.Author.ID}
		err := users.FindOne(db.Ctx(), filter).Decode(&user)

		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Utilizador não existente, registe-se com `!register`")
		} else {
			_, err := users.UpdateOne(db.Ctx(), filter, bson.D{
				{Key: "$set", Value: bson.M{"city": city}},
			})
			if err != nil {
				session.ChannelMessageSend(command.ChannelID, "Erro ao alterar a sua cidade default")
			} else {
				session.ChannelMessageSend(command.ChannelID, "Cidade alterada com sucesso")
			}
		}
	} else {
		session.ChannelMessageSend(command.ChannelID, "Escolha uma localizacão, exemplo: `!mycity Mafamude`")
	}
}
