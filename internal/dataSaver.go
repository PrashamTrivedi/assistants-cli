package internal

type Assistants []Assistant

type Assistant struct {
	ID           string `json:"id,omitempty"`
	Prompt       string `json:"prompt,omitempty"`
	Name         string `json:"name,omitempty"`
	DefaultModel string `json:"default_model,omitempty"`
}

type ChatData struct {
	ID        string    `json:"id,omitempty"`
	Assistant string    `json:"assistant,omitempty"`
	Messages  []Message `json:"messages,omitempty"`
}

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

type AssistantStore interface {
	WriteAssistants(Assistants) error
	CreateAssistant(Assistant) (Assistants, error)
	UpdateAssistant(Assistant) (Assistants, error)
	RemoveAssistant(Assistant) (Assistants, error)
	ReadAssistants() (Assistants, error)
}

type ChatStore interface {
	WriteChats([]ChatData) error
	CreateChat(ChatData) ([]ChatData, error)
	UpdateChat(ChatData) ([]ChatData, error)
	RemoveChat(ChatData) ([]ChatData, error)
	ReadChats() ([]ChatData, error)
}
