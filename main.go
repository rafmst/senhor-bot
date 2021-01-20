package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	discord, err := discordgo.New("Bot " + "authentication token")
}
