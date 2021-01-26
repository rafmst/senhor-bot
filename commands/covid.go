package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	thousands "github.com/floscodes/golang-thousands"
)

// Country is the covid type of country
type Country struct {
	Name        string      `json:"country"`
	CountryInfo countryInfo `json:"countryInfo"`
	Cases       int         `json:"todayCases"`
	Deaths      int         `json:"todayDeaths"`
	Updated     int64       `json:"updated"`
}

type countryInfo struct {
	Flag string `json:"flag"`
}

// HandleCovid handles "!covid" command
func HandleCovid(session *discordgo.Session, command *discordgo.MessageCreate) {
	showCountryInfo("Portugal", session, command)
	showCountryInfo("Norway", session, command)
	showCountryInfo("France", session, command)
}

func showCountryInfo(country string, session *discordgo.Session, command *discordgo.MessageCreate) {
	resp, err := http.Get("https://corona.lmao.ninja/v2/countries/" + country)

	if err != nil {
		session.ChannelMessageSend(command.ChannelID, "Error: Covid Server seems to be down")
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			session.ChannelMessageSend(command.ChannelID, "Error: Wasn't able to read response")
		} else {
			var country Country
			err := json.Unmarshal(body, &country)
			if err != nil {
				session.ChannelMessageSend(command.ChannelID, "Error: Check Covid API  response format")
			} else {
				thumbnail := discordgo.MessageEmbedThumbnail{
					URL:    country.CountryInfo.Flag,
					Width:  36,
					Height: 36,
				}

				cases := readableThousands(country.Cases)
				deaths := readableThousands(country.Deaths)
				updated := time.Unix(country.Updated/1000, 0)

				var message = discordgo.MessageEmbed{
					Title: country.Name,
					Description: `**` + cases + `** Casos
					**` + deaths + `** Mortes
					
					*Última actualização: ` + updated.Format("2006-01-02 15:04") + `*`,
					Thumbnail: &thumbnail,
				}

				session.ChannelMessageSendEmbed(command.ChannelID, &message)
			}
		}
	}
}

func readableThousands(number int) string {
	return strings.Replace(thousands.Separate(strconv.Itoa(number), "de"), ".", " ", 2)
}
