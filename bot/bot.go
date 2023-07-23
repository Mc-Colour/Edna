package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	BotToken string
	//OpenAiToken string
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
	discord.Open()
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
	switch {

	//case strings.Contains(message.Content, "!Edna"):
	//botResponse := GptResponse(message.Content)
	//discord.ChannelMessageSend(message.ChannelID, botResponse)
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hi there!")
	}

}

// GPT-3 stuff
//func GptResponse(prompt string) string {
//	client := openai.NewClient(OpenAiToken)
//	resp, err := client.CreateChatCompletion(
//		context.Background(),
//		openai.ChatCompletionRequest{
//			Model: openai.GPT3Dot5Turbo,
//			Messages: []openai.ChatCompletionMessage{
//				{
//					Role:    openai.ChatMessageRoleUser,
//					Content: prompt,
//				},
//			},
//		},
//	)
//
//	if err != nil {
//		fmt.Printf("ChatCompletion error: %v\n", err)
//		return "Chat completion error :("
//	}
//	fmt.Printf(resp.Choices[0].Message.Content)
//	return resp.Choices[0].Message.Content
//}
