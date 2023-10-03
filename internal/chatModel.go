package internal

import (
	"context"
	"encoding/json"
	"fmt"

	"log/slog"

	"github.com/oklog/ulid/v2"
	openai "github.com/sashabaranov/go-openai"
)

type ApiKey string
type SerpKey string

type Chat struct {
	chatId     string
	assistant  *Assistant
	messages   []openai.ChatCompletionMessage
	chatClient *openai.Client
	functions  []openai.FunctionDefinition
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
	chat.functions = getFunctionDefinitions(assistant)
	chat.chatClient = client

	return chat, nil
}

func getFunctionDefinitions(assistant Assistant) []openai.FunctionDefinition {
	functions := make([]openai.FunctionDefinition, 0)

	if assistant.AllowSearch {
		functions = append(functions, openai.FunctionDefinition{
			Name:        "search",
			Description: "Search the web for the query",
			Parameters: json.RawMessage(`{
    "type": "object",
    "properties": {
        "query": {
            "type": "string"
        }
    },
    "required": [
        "query"
    ]
}`)})
	}
	if assistant.AllowCommands {
		functions = append(functions, openai.FunctionDefinition{
			Name:        "command",
			Description: "Execute a command on the machine with the arguments",
			Parameters: json.RawMessage(`{
	"type": "object",
	"properties": {
		"command": {
			"type": "string"
		},
		"arguments": {
			"type": "string"
		}
	},
	"required": [
		"command"
	]}`)})

	}

	if assistant.AllowFileReading {
		functions = append(functions, openai.FunctionDefinition{
			Name:        "read_file",
			Description: "Read a file from the machine",
			Parameters: json.RawMessage(`{
	"type": "object",
	"properties": {
		"file": {
			"type": "string"
		}
	},
	"required": [
		"file"
	]}`)})
	}

	return functions
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

func (c *Chat) Start(message string, serpApiKey SerpKey, chatStore ChatStore) (string, error) {
	fmt.Println("Starting chat...")
	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
	resp, err := c.chatClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     c.assistant.DefaultModel,
			Messages:  c.messages,
			Functions: c.functions,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	responseMessage := resp.Choices[0].Message
	// Arguments is a stringified json, convert it into a map
	functionMessages, err := handleFunctionCall(responseMessage, resp, serpApiKey, c)

	if err != nil {
		fmt.Println("Error handling function call: ", err.Error())
		return "", err
	}
	isFunctionHandled := err == nil && len(functionMessages) > 0
	if isFunctionHandled {
		c.messages = append(c.messages, functionMessages...)
	} else {
		c.messages = append(c.messages, resp.Choices[0].Message)
		fmt.Println(resp.Choices[0].Message.Content)
	}

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

func handleFunctionCall(responseMessage openai.ChatCompletionMessage, resp openai.ChatCompletionResponse, serpApiKey SerpKey, c *Chat) ([]openai.ChatCompletionMessage, error) {

	functionMessages := make([]openai.ChatCompletionMessage, 0)
	if responseMessage.FunctionCall != nil && responseMessage.FunctionCall.Name != "" {
		functionMessage := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleFunction,
			Content: "",
		}
		slog.Info("Function call: ", "Name", resp.Choices[0].Message.FunctionCall.Name)
		slog.Info("Function call: ", "Args", resp.Choices[0].Message.FunctionCall.Arguments)
		functionMessage.Name = resp.Choices[0].Message.FunctionCall.Name

		var arguments map[string]interface{}
		err := json.Unmarshal([]byte(resp.Choices[0].Message.FunctionCall.Arguments), &arguments)
		if err != nil {
			fmt.Println("Error parsing arguments: ", err.Error())

			functionMessage.Content = "I had some trouble understanding your request. It can be the query, or it can be my understanding. "
		}

		if resp.Choices[0].Message.FunctionCall.Name == "search" {
			query, ok := arguments["query"].(string)
			if !ok {
				functionMessage.Content = "The query I got is not a string. "
			}
			results, err := Search(query, string(serpApiKey))
			if err != nil {
				functionMessage.Content = "I had some trouble searching for your query. "
			}
			slog.Debug("Results: ", "Result", results)
			resultOutput := fmt.Sprintf("{\"snippet\":\"%s\",\"link\":\"(%s)\"}",
				results["snippet"].(string), results["link"].(string))
			functionMessage.Content = resultOutput
		}
		if resp.Choices[0].Message.FunctionCall.Name == "command" {
			command, ok := arguments["command"].(string)
			if !ok {
				functionMessage.Content = "The command I got is not a string, please try again"
			}
			args, ok := arguments["arguments"].(string)
			if !ok {
				functionMessage.Content = "The arguments I got is not a string, please try again"
			}
			slog.Info("Command", "Command", command, "Arguments", args)
			output, err := RunCommand(command, args)
			slog.Info("Command", "Output", output)
			if err != nil {
				functionMessage.Content = fmt.Sprintf("The function encountered this error: '%s', State the error and then explain this", err.Error())
			} else if output == "" {
				functionMessage.Content = "I did not get any output from your command. "
			} else {
				functionMessage.Content = output
			}
		}
		if resp.Choices[0].Message.FunctionCall.Name == "read_file" {
			file, ok := arguments["file"].(string)
			if !ok {
				functionMessage.Content = "The file I got is not a string, please try again"
			}
			slog.Info("File", "File", file)

			output, err := Readfile(file)
			slog.Info("File", "Output", output)
			if err != nil {
				slog.Error("File", "Error", err.Error())
				if err.Error() == "file does not exist" {
					functionMessage.Content = fmt.Sprintf("The file %s does not exist. ", file)
				} else {
					functionMessage.Content = fmt.Sprintf("The function encountered this error: '%s'", err.Error())
				}
			} else {
				functionMessage.Content = output
			}
		}
		functionMessages = append(functionMessages, functionMessage)
		messages := append(c.messages, functionMessage)
		chatRequest := openai.ChatCompletionRequest{
			Model:     c.assistant.DefaultModel,
			Messages:  messages,
			Functions: c.functions,
		}
		afterFuncResponse, err := c.chatClient.CreateChatCompletion(
			context.Background(),
			chatRequest,
		)
		if err != nil {
			slog.Error("ChatCompletion", "error", err)
			return nil, err
		}
		functionMessages = append(functionMessages, afterFuncResponse.Choices[0].Message)
		fmt.Println(afterFuncResponse.Choices[0].Message.Content)
	}
	return functionMessages, nil
}

func (c *Chat) Continue(message string, serpApiKey SerpKey, chatStore ChatStore, assistantsStore AssistantStore) error {

	userMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	}
	for index, message := range c.messages {
		if message.Role == openai.ChatMessageRoleFunction && message.Name == "" {
			message.Name = "function"
		}
		c.messages[index] = message
	}
	c.messages = append(c.messages, userMessage)
	resp, err := c.chatClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     c.assistant.DefaultModel,
			Messages:  c.messages,
			Functions: c.functions,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}
	responseMessage := resp.Choices[0].Message
	// Arguments is a stringified json, convert it into a map
	functionMessages, err := handleFunctionCall(responseMessage, resp, serpApiKey, c)

	if err != nil {
		fmt.Println("Error handling function call: ", err.Error())
		return err
	}
	isFunctionHandled := err != nil && len(functionMessages) > 0
	if isFunctionHandled {
		c.messages = append(c.messages, functionMessages...)
	} else {
		c.messages = append(c.messages, resp.Choices[0].Message)
		fmt.Println(resp.Choices[0].Message.Content)
	}

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
			Name:    message.Name,
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
