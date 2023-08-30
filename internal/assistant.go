package internal

import (
	"fmt"
	"os"

	"github.com/oklog/ulid/v2"
)

func NewAssistant(name, prompt, model string, assistantStore AssistantStore) *Assistants {
	assistant := Assistant{Name: name, Prompt: prompt, DefaultModel: model}
	assistant.ID = ulid.Make().String()

	assistants, err := assistantStore.CreateAssistant(assistant)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return &assistants
}

func RemoveAssistant(name string, assistantStore AssistantStore) (Assistants, error) {
	return assistantStore.RemoveAssistant(Assistant{Name: name})
}

func UpdateAssistant(name, prompt, model string, assistantStore AssistantStore) (Assistants, error) {
	return assistantStore.UpdateAssistant(Assistant{Name: name, Prompt: prompt, DefaultModel: model})
}

func FindAssistant(name string, assistantStore AssistantStore) (*Assistant, error) {
	assistants, err := assistantStore.ReadAssistants()

	if err != nil {
		return nil, err
	}
	for _, assistant := range assistants {
		if assistant.Name == name {
			return &assistant, nil
		}
	}
	return nil, fmt.Errorf("assistant %s not found", name)
}
