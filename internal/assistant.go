package internal

import (
	"fmt"
	"os"
)

func NewAssistant(name, prompt, model string, assistantStore AssistantStore) *Assistants {
	assistants, err := assistantStore.ReadAssistants()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, assistant := range assistants {
		if assistant.Name == name {
			fmt.Println("Assistant already exists")
			os.Exit(1)
		}
	}
	assistants = append(assistants, Assistant{Prompt: prompt, Name: name})
	assistantStore.WriteAssistants(assistants)
	return &assistants
}

func RemoveAssistant(name string, assistantStore AssistantStore) error {
	assistants, err := assistantStore.ReadAssistants()

	if err != nil {
		return err
	}
	for i, assistant := range assistants {
		if assistant.Name == name {
			assistants = append(assistants[:i], assistants[i+1:]...)
			break
		}
	}
	assistantStore.WriteAssistants(assistants)

	// implementation to remove assistant
	return nil
}

func UpdateAssistant(name, prompt, model string, assistantStore AssistantStore) error {
	assistants, err := assistantStore.ReadAssistants()

	if err != nil {
		return err
	}
	for i, assistant := range assistants {
		if assistant.Name == name {
			if prompt != "" {
				assistants[i].Prompt = prompt
			}
			if model != "" {
				assistants[i].DefaultModel = model
			}
			break
		}
	}
	assistantStore.WriteAssistants(assistants)

	// implementation to update assistant
	return nil
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
