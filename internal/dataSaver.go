package internal

type Assistants []Assistant

type Assistant struct {
	ID               string `json:"id,omitempty"`
	Prompt           string `json:"prompt,omitempty"`
	Name             string `json:"name,omitempty"`
	DefaultModel     string `json:"default_model,omitempty"`
	CreatedOn        int64  `json:"created_on,omitempty"`
	UpdatedOn        int64  `json:"updated_on,omitempty"`
	AllowSearch      bool   `json:"allow_search,omitempty"`
	AllowCommands    bool   `json:"allow_commands,omitempty"`
	AllowFileReading bool   `json:"allow_file_reading,omitempty"`
}

type ChatData struct {
	ID          string      `json:"id,omitempty"`
	AssistantId string      `json:"assistant_id,omitempty"`
	Messages    []Message   `json:"messages,omitempty"`
	CreatedOn   int64       `json:"created_on,omitempty"`
	UpdatedOn   int64       `json:"updated_on,omitempty"`
	Functions   []Functions `json:"functions,omitempty"`
}

type Functions struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Parameters  any    `json:"parameters"`
}

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
}

type AssistantStore interface {
	WriteAssistants(Assistants) error
	CreateAssistant(Assistant) (Assistants, error)
	UpdateAssistant(Assistant) (Assistants, error)
	RemoveAssistant(Assistant) (Assistants, error)
	FindAssistant(string) (*Assistant, error)
	ReadAssistants() (Assistants, error)
}

type ChatStore interface {
	WriteChats([]ChatData) error
	CreateChat(ChatData) ([]ChatData, error)
	AddNewChatMessage(string, []Message) ([]ChatData, error)
	GetChat(string) (ChatData, error)
	RemoveChat(ChatData) ([]ChatData, error)
	ReadChats() ([]ChatData, error)
}
