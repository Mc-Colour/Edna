package main

import (
	"Edna/bot"
	"log"
	"os"
)

func main() {
	// Load environment variables
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	openAiToken, ok := os.LookupEnv("OPENAI_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	// Start the bot
	bot.BotToken = botToken
	bot.OpenAiToken = openAiToken
	bot.Run()
}
