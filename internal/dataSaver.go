package internal

type Assistants []Assistant

type Assistant struct {
	Prompt       string `json:"prompt,omitempty"`
	Name         string `json:"name,omitempty"`
	DefaultModel string `json:"default_model,omitempty"`
}
type AssistantStore interface {
	WriteAssistants(Assistants) error

	ReadAssistants() (Assistants, error)
}

// type ChatSaver interface {
// 	WriteChats([]Chat) error
// 	ReadChats() ([]Chat, error)
// }
