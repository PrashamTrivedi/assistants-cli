package filesavers

import (
	"assistants-cli/internal"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type AssistantFileStore struct {
	AssistantFilePath string
}

func GetDefaultAssistantFilePath() string {
	home, _ := os.UserHomeDir()
	assistantsPath := filepath.Join(home, ".assistants", "assistants.json")
	return assistantsPath
}

func NewAssistantFileStore(assistantFilePath string) *AssistantFileStore {
	if assistantFilePath == "" {
		assistantFilePath = GetDefaultAssistantFilePath()
		internal.WriteConfig(map[string]string{internal.AssistantFilePath: assistantFilePath})
	}
	// Check if file exists on assistantFilePath, if not create it
	if _, err := os.Stat(assistantFilePath); os.IsNotExist(err) {
		file, err := os.Create(assistantFilePath)
		defer file.Close()
		if err != nil {
			panic(err)
		}
		// Write empty array to file
		emptyAssistants := internal.Assistants{}
		if err := json.NewEncoder(file).Encode(emptyAssistants); err != nil {
			panic(err)
		}
	}
	return &AssistantFileStore{AssistantFilePath: assistantFilePath}
}

func (f *AssistantFileStore) WriteAssistants(assistants internal.Assistants) error {
	// Write assistants to embedded file
	assistantFile, error := os.OpenFile(f.AssistantFilePath, os.O_CREATE|os.O_WRONLY, fs.FileMode(os.O_RDWR))
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

func (f *AssistantFileStore) FindAssistant(assistantId string) (*internal.Assistant, error) {

	assistants, err := f.getAssistants()

	if err != nil {
		return nil, err
	}
	for _, assistant := range assistants {
		if assistant.ID == assistantId {
			return &assistant, nil
		}
	}
	return nil, nil
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
			assistants[i].AllowSearch = assistant.AllowSearch
			assistants[i].AllowCommands = assistant.AllowCommands
			assistants[i].AllowFileReading = assistant.AllowFileReading
			assistant.UpdatedOn = time.Now().UnixMilli()
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
