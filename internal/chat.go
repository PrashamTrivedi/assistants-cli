package internal

import (
	"context"
	"fmt"

	"github.com/oklog/ulid/v2"
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

func ListChat(chatStore ChatStore) ([]ChatData, error) {
	chats, err := chatStore.ReadChats()
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (c *Chat) Start(message string, chatStore ChatStore) error {
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

	chatData := ChatData{
		ID:          ulid.Make().String(),
		AssistantId: c.assistant.ID,
		Messages:    c.convertMessageForStorage(),
	}
	_, err = chatStore.CreateChat(chatData)
	if err != nil {
		return err
	}

	return nil
}

func (c *Chat) Continue(chatId, message string, chatStore ChatStore, assistantsStore AssistantStore) error {
	chatData, err := chatStore.GetChat(chatId)
	if err != nil {
		return err
	}
	assistant, err := assistantsStore.FindAssistant(chatData.AssistantId)
	if err != nil {
		return err
	}
	c.assistant = assistant
	c.convertMessageFromStorage(chatData)
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

	_, err = chatStore.AddNewChatMessage(chatId, message)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chat) convertMessageForStorage() []Message {
	var messages []Message
	for _, message := range c.messages {
		messages = append(messages, Message{
			Role:    message.Role,
			Content: message.Content,
		})
	}
	return messages
}

func (c *Chat) convertMessageFromStorage(chatData ChatData) error {
	for _, message := range chatData.Messages {
		c.messages = append(c.messages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}
	return nil
}
