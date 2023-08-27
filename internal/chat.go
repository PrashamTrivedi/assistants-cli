package internal

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type ApiKey string

type Chat struct {
	assistant  *Assistant
	messages   []openai.ChatCompletionMessage
	chatClient *openai.Client
}

func NewChat(apiKey ApiKey, assistant Assistant) (*Chat, error) {
	chat := &Chat{}
	client := openai.NewClient(string(apiKey))
	chat.messages = []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: assistant.Prompt,
		},
	}
	chat.chatClient = client

	return chat, nil
}

func (c *Chat) Start(message string) error {
	fmt.Println("Starting chat...")
	resp, err := c.chatClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    c.assistant.DefaultModel,
			Messages: c.messages,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return nil
}

func (c *Chat) Continue(message string) error {
	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
	resp, err := c.chatClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    c.assistant.DefaultModel,
			Messages: c.messages,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return nil
}

func getInput() string {
	var input string
	fmt.Print("> ")
	fmt.Scanln(&input)
	return input
}
