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
**!covid**: Mostra o número de casos e mortes para Portugal e Noruega;
**!register**: Regista o utilizador na base de dados;
**!mycity**: Regista o a tua cidade default. Exemplo: ` + "`!mycity Mafamude`" + `;
**!unregister**: Apagar os seus dados da base de dados;
**!users**: Lista de utizadores e seus detalhes;
**!mowgli**: Mostra uma imagem aleatória do Mowgli;
**!napoleao**: Mostra uma imagem aleatória do Napoleão;
**!safira**: Mostra uma imagem aleatória da Safira;
`,
	}

	session.ChannelMessageSendEmbed(command.ChannelID, &message)
}
