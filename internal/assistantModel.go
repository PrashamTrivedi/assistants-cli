package internal

import (
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewAssistant(name, prompt, model string, allowSearch, allowCommands, allowFileReading bool, assistantStore AssistantStore) (*Assistants, error) {
	assistant := Assistant{Name: name, Prompt: prompt, DefaultModel: model, AllowSearch: allowSearch, AllowCommands: allowCommands, AllowFileReading: allowFileReading}
	assistant.ID = ulid.Make().String()
	assistant.CreatedOn = time.Now().UnixMilli()

	assistants, err := assistantStore.CreateAssistant(assistant)

	return &assistants, err
}

func RemoveAssistant(id string, assistantStore AssistantStore) (Assistants, error) {
	return assistantStore.RemoveAssistant(Assistant{ID: id})
}

func UpdateAssistant(id, name, prompt, model string, allowSearch, allowCommands, allowFileReading bool, assistantStore AssistantStore) (Assistants, error) {
	return assistantStore.UpdateAssistant(Assistant{
		ID: id, Name: name, Prompt: prompt, DefaultModel: model, AllowSearch: allowSearch, AllowCommands: allowCommands, AllowFileReading: allowFileReading})
}

func ReadAssistants(assistantStore AssistantStore) ([]Assistant, error) {
	return assistantStore.ReadAssistants()
}

func FindAssistant(assistantId string, assistantStore AssistantStore) (*Assistant, error) {
	assistants, err := assistantStore.ReadAssistants()

	if err != nil {
		return nil, err
	}
	for _, assistant := range assistants {
		if assistant.ID == assistantId {
			return &assistant, nil
		}
	}
	return nil, fmt.Errorf("assistant %s not found", assistantId)
}
