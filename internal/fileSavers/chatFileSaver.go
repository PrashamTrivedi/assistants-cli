package filesavers

import (
	"assistants-cli/internal"
	"encoding/json"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
)

type ChatFileStore struct {
	ChatFilePath string
}

func NewChatFileStore(chatFilePath string) *ChatFileStore {
	if _, err := os.Stat(chatFilePath); os.IsNotExist(err) {
		file, err := os.Create(chatFilePath)
		defer file.Close()
		if err != nil {
			panic(err)
		}
		// Write empty array to file
		emptyChats := []internal.ChatData{}
		if err := json.NewEncoder(file).Encode(emptyChats); err != nil {
			panic(err)
		}

	}
	return &ChatFileStore{ChatFilePath: chatFilePath}
}

func (f *ChatFileStore) WriteChats(chats []internal.ChatData) error {
	// Write chats to embedded file
	chatFile, error := os.Open(f.ChatFilePath)
	if error != nil {
		return error
	}
	if err := json.NewEncoder(chatFile).Encode(chats); err != nil {
		return err
	}
	return nil
}

func (f *ChatFileStore) getChats() ([]internal.ChatData, error) {
	chatFile, error := os.Open(f.ChatFilePath)
	if error != nil {
		return nil, error
	}
	var chats []internal.ChatData
	if err := json.NewDecoder(chatFile).Decode(&chats); err != nil {
		return nil, err
	}
	return chats, nil

}

func (f *ChatFileStore) CreateChat(chat internal.ChatData) ([]internal.ChatData, error) {

	chat.CreatedOn = time.Now().UnixMilli()
	chats, err := f.getChats()
	if err != nil {
		return nil, err
	}
	chats = append(chats, chat)
	err = f.WriteChats(chats)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (f *ChatFileStore) AddNewChatMessage(chatId string, message string) ([]internal.ChatData, error) {
	chats, err := f.getChats()
	if err != nil {
		return nil, err
	}
	for i, chatFromStore := range chats {
		if chatId == chatFromStore.ID {
			chats[i].Messages = append(chats[i].Messages, internal.Message{Role: openai.ChatMessageRoleUser, Content: message})
			chats[i].UpdatedOn = time.Now().UnixMilli()
		}
	}
	err = f.WriteChats(chats)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (f *ChatFileStore) GetChat(id string) (internal.ChatData, error) {
	chats, err := f.getChats()
	if err != nil {
		return internal.ChatData{}, err
	}
	for _, chat := range chats {
		if chat.ID == id {
			return chat, nil
		}
	}
	return internal.ChatData{}, nil
}

func (f *ChatFileStore) RemoveChat(chat internal.ChatData) ([]internal.ChatData, error) {
	chats, err := f.getChats()
	if err != nil {
		return nil, err
	}
	for i, chatFromStore := range chats {
		if chat.ID == chatFromStore.ID {
			chats = append(chats[:i], chats[i+1:]...)
		}
	}
	err = f.WriteChats(chats)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (f *ChatFileStore) ReadChats() ([]internal.ChatData, error) {
	// Read chats from embedded file
	return f.getChats()
}
