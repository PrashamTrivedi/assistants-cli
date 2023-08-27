package internal

import (
	"encoding/json"
	"os"
)

type FileStore struct {
	AssistantFilePath string
}

func NewFileStore(assistantFilePath string) *FileStore {
	return &FileStore{AssistantFilePath: assistantFilePath}
}

func (f *FileStore) WriteAssistants(assistants Assistants) error {
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

func (f *FileStore) ReadAssistants() (Assistants, error) {
	// Read assistants from embedded file
	assistantFile, error := os.Open(f.AssistantFilePath)
	if error != nil {
		return nil, error
	}
	var assistants Assistants
	if err := json.NewDecoder(assistantFile).Decode(&assistants); err != nil {
		return nil, err
	}

	return assistants, nil
}
