package internal

import (
	"context"
	"fmt"

	"github.com/oklog/ulid/v2"
	openai "github.com/sashabaranov/go-openai"
)

type ApiKey string

type Chat struct {
	chatId     string
	assistant  *Assistant
	messages   []openai.ChatCompletionMessage
	chatClient *openai.Client
}

func NewChat(apiKey ApiKey, assistant Assistant) (*Chat, error) {
	chat := &Chat{}
	client := openai.NewClient(string(apiKey))
	chat.assistant = &assistant
	chat.messages = []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: assistant.Prompt,
			Name:    assistant.Name,
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

func GetChat(chatId string, apiKey ApiKey, chatStore ChatStore, assistantStore AssistantStore) (*Chat, error) {
	chatData, err := chatStore.GetChat(chatId)
	if err != nil {
		return nil, err
	}
	chat := Chat{}
	chat.convertMessageFromStorage(chatData)
	chat.assistant, err = FindAssistant(chatData.AssistantId, assistantStore)
	if err != nil {
		return nil, err
	}
	chat.chatClient = openai.NewClient(string(apiKey))
	chat.chatId = chatData.ID

	return &chat, nil
}

func (c *Chat) Start(message string, chatStore ChatStore) (string, error) {
	fmt.Println("Starting chat...")
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
		return "", err
	}

	fmt.Println(resp.Choices[0].Message.Content)
	c.messages = append(c.messages, resp.Choices[0].Message)

	chatData := ChatData{
		ID:          ulid.Make().String(),
		AssistantId: c.assistant.ID,
		Messages:    c.convertMessageForStorage(),
	}

	_, err = chatStore.CreateChat(chatData)
	if err != nil {
		return "", err
	}

	return chatData.ID, nil
}

func (c *Chat) Continue(message string, chatStore ChatStore, assistantsStore AssistantStore) error {

	userMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	}
	c.messages = append(c.messages, userMessage)
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
	c.messages = append(c.messages, resp.Choices[0].Message)

	messagesToSend := []Message{{
		Role:    userMessage.Role,
		Content: userMessage.Content,
	}, {
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}}
	_, err = chatStore.AddNewChatMessage(c.chatId, messagesToSend)
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
