package commands

import (
	"github.com/bwmarrin/discordgo"
)

// HandleHelp handles "!help" command
func HandleHelp(session *discordgo.Session, command *discordgo.MessageCreate) {
	var message = discordgo.MessageEmbed{
		Title: "Comandos disponíveis",
		Description: `**!help**: Parece-me que já sabes usar este commando;
		**!dog**: Mostra uma imagem aleatória de um cão;
		**!cat**: Mostra uma imagem aleatória de um gato;
		**!fox**: Mostra uma imagem aleatória de uma raposa;
		**!weather** (**!w**): Mostra o tempo para o sítio escolhido;
		`,
	}

	session.ChannelMessageSendEmbed(command.ChannelID, &message)
}
