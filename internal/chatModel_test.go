package internal

import "testing"

func NewChatStore() ChatStore {
	return &InMemoryChatStore{}
}

// Implement ChatStore interface
type InMemoryChatStore struct {
	chats []ChatData
}

func (i *InMemoryChatStore) WriteChats(chats []ChatData) error {
	i.chats = chats
	return nil
}

func (i *InMemoryChatStore) CreateChat(chat ChatData) ([]ChatData, error) {
	i.chats = append(i.chats, chat)
	return i.chats, nil
}

func (i *InMemoryChatStore) AddNewChatMessage(chatId string, messages []Message) ([]ChatData, error) {
	for index, chat := range i.chats {
		if chat.ID == chatId {
			i.chats[index].Messages = append(i.chats[index].Messages, messages...)
		}
	}
	return i.chats, nil
}

func (i *InMemoryChatStore) GetChat(chatId string) (ChatData, error) {
	for _, chat := range i.chats {
		if chat.ID == chatId {
			return chat, nil
		}
	}
	return ChatData{}, nil
}

func (i *InMemoryChatStore) RemoveChat(chat ChatData) ([]ChatData, error) {
	for index, chatFromStore := range i.chats {
		if chatFromStore.ID == chat.ID {
			i.chats = append(i.chats[:index], i.chats[index+1:]...)
		}
	}
	return i.chats, nil
}

func (i *InMemoryChatStore) ReadChats() ([]ChatData, error) {
	return i.chats, nil
}

func TestNewChat(t *testing.T) {
	// Initialize dependencies
	assistant := Assistant{ID: "testAssistant"}
	apiKey := ApiKey("testKey")

	// Test the NewChat function
	chat, err := NewChat(apiKey, assistant)

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if chat == nil {
		t.Fatalf("Expected non-nil Chat, got nil")
	}
	if *chat.assistant != assistant {
		t.Errorf("Expected assistant to be %v, got %v", assistant, *chat.assistant)
	}
}

func TestListChat(t *testing.T) {
	// Initialize MockChatStore with some dummy data
	mockChatStore := NewChatStore()
	mockAssistantStore := NewAssistantStore()

	assistants, error := NewAssistant("testAssistant", "Say Hello", "davinci", mockAssistantStore)
	if error != nil {
		t.Fatalf("Expected nil error, got %v", error)
	}
	assistant := (*assistants)[0]

	chat, error := NewChat("testKey", assistant)

	if error != nil {
		t.Fatalf("Expected nil error, got %v", error)
	}

	_, error = chat.Start("Hello", mockChatStore)
	if error != nil {
		t.Fatalf("Expected nil error, got %v", error)
	}
	error = chat.Continue("Hello", mockChatStore, mockAssistantStore)
	if error != nil {
		t.Fatalf("Expected nil error, got %v", error)
	}

	// Test the ListChat function
	chats, err := ListChat(mockChatStore)

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if chats == nil {
		t.Fatalf("Expected non-nil chats, got nil")
	}
	if len(chats) != 2 {
		t.Errorf("Expected number of chats to be 2, got %d", len(chats))
	}
}

func TestGetChat(t *testing.T) {
	// Initialize MockChatStore with some dummy data
	mockChatStore := NewChatStore()
	assistantStore := NewAssistantStore()
	apiKey := ApiKey("testKey")

	// Test the GetChat function
	chat, err := GetChat("chat1", apiKey, mockChatStore, assistantStore)

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if chat == nil {
		t.Fatalf("Expected non-nil chat, got nil")
	}
	if chat.chatId != "chat1" {
		t.Errorf("Expected chat ID to be 'chat1', got %s", chat.chatId)
	}

	// Test with a non-existing chat ID
	_, err = GetChat("nonexistent", apiKey, mockChatStore, assistantStore)
	if err == nil {
		t.Fatalf("Expected an error for non-existing chat ID, got nil")
	}
}

func TestStart(t *testing.T) {
	// Initialize MockChatStore with some dummy data
	mockChatStore := NewChatStore()

	// Initialize a Chat object
	chat := &Chat{
		chatId: "chat1",
	}

	// Test the Start function
	chatID, err := chat.Start("Hello", mockChatStore)

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if chatID != "chat1" {
		t.Errorf("Expected chat ID to be 'chat1', got %s", chatID)
	}
}
