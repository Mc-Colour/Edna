package bot

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	BotToken    string
	OpenAiToken string
)

func Run() {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	// Add event handler
	discord.AddHandler(newMessage)

	// Open session
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer discord.Close()

	// Run until code is terminated
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot message
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	//switch {
	//
	//case strings.Contains(message.Content, "Edna"):
	//	botResponse := GptResponse(message.Content)
	//	discord.ChannelMessageSend(message.ChannelID, botResponse)
	//}
	if strings.Contains(message.Content, "Edna") || strings.Contains(message.Content, "edna") {
		botResponse := GptResponse(message.Content)
		discord.ChannelMessageSend(message.ChannelID, botResponse)
	}

}

// GPT-3 stuff
func GptResponse(prompt string) string {
	client := openai.NewClient(OpenAiToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "I need more money honey"
	}
	return resp.Choices[0].Message.Content
}
