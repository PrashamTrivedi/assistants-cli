package filesavers

import (
	"assistants-cli/internal"
	"encoding/json"
	"os"
)

type AssistantFileStore struct {
	AssistantFilePath string
}

func NewAssistantFileStore(assistantFilePath string) *AssistantFileStore {
	return &AssistantFileStore{AssistantFilePath: assistantFilePath}
}

func (f *AssistantFileStore) WriteAssistants(assistants internal.Assistants) error {
	// Write assistants to embedded file
	assistantFile, error := os.Open(f.AssistantFilePath)
	if error != nil {
		return error
	}
	if err := json.NewEncoder(assistantFile).Encode(assistants); err != nil {
		return err
	}
	return nil
}

func (f *AssistantFileStore) getAssistants() (internal.Assistants, error) {
	assistantFile, error := os.Open(f.AssistantFilePath)
	if error != nil {
		return nil, error
	}
	var assistants internal.Assistants
	if err := json.NewDecoder(assistantFile).Decode(&assistants); err != nil {
		return nil, err
	}
	return assistants, nil

}

func (f *AssistantFileStore) CreateAssistant(assistant internal.Assistant) (internal.Assistants, error) {

	assistants, err := f.getAssistants()
	if err != nil {
		return nil, err
	}
	assistants = append(assistants, assistant)
	err = f.WriteAssistants(assistants)
	if err != nil {
		return nil, err
	}
	return assistants, nil
}

func (f *AssistantFileStore) UpdateAssistant(assistant internal.Assistant) (internal.Assistants, error) {
	assistants, err := f.getAssistants()

	if err != nil {
		return nil, err
	}
	for i, assistantFromStore := range assistants {
		if assistant.ID == assistantFromStore.ID {
			if assistant.Prompt != "" {
				assistants[i].Prompt = assistant.Prompt
			}
			if assistant.DefaultModel != "" {
				assistants[i].DefaultModel = assistant.DefaultModel
			}
			break
		}
	}
	err = f.WriteAssistants(assistants)
	if err != nil {
		return nil, err
	}
	return assistants, nil
}

func (f *AssistantFileStore) ReadAssistants() (internal.Assistants, error) {
	// Read assistants from embedded file
	return f.getAssistants()
}

func (f *AssistantFileStore) RemoveAssistant(assistant internal.Assistant) (internal.Assistants, error) {
	assistants, err := f.getAssistants()

	if err != nil {
		return nil, err
	}
	for i, assistantFromStore := range assistants {
		if assistant.ID == assistantFromStore.ID {
			assistants = append(assistants[:i], assistants[i+1:]...)
			break
		}
	}
	err = f.WriteAssistants(assistants)
	if err != nil {
		return nil, err
	}
	return assistants, nil
}
