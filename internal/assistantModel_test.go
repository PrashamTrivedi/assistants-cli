package internal

import (
	"fmt"
	"testing"
)

// Give me in memory assistant store that implements AssistantStore interface
func NewAssistantStore() AssistantStore {
	return &InMemoryAssistantStore{}
}

// Implement AssistantStore interface
type InMemoryAssistantStore struct {
	assistants Assistants
}

func (i *InMemoryAssistantStore) WriteAssistants(assistants Assistants) error {
	i.assistants = assistants
	return nil
}

func (i *InMemoryAssistantStore) CreateAssistant(assistant Assistant) (Assistants, error) {
	if assistant.DefaultModel == "ErrorModel" {
		return nil, fmt.Errorf("An error is thrown here")
	}
	i.assistants = append(i.assistants, assistant)
	return i.assistants, nil
}

func (i *InMemoryAssistantStore) UpdateAssistant(assistant Assistant) (Assistants, error) {
	for index, assistantFromStore := range i.assistants {
		if assistantFromStore.ID == assistant.ID {
			i.assistants[index] = assistant
		}
	}
	return i.assistants, nil
}

func (i *InMemoryAssistantStore) FindAssistant(id string) (*Assistant, error) {
	for _, assistant := range i.assistants {
		if assistant.ID == id {
			return &assistant, nil
		}
	}
	return nil, nil
}

func (i *InMemoryAssistantStore) ReadAssistants() (Assistants, error) {
	for _, assistant := range i.assistants {
		if assistant.Name == "error" {
			return nil, fmt.Errorf("An error is thrown here")
		}
	}
	return i.assistants, nil
}

func (i *InMemoryAssistantStore) RemoveAssistant(assistant Assistant) (Assistants, error) {
	for index, assistantFromStore := range i.assistants {
		if assistantFromStore.ID == assistant.ID {
			i.assistants = append(i.assistants[:index], i.assistants[index+1:]...)
		}
	}
	return i.assistants, nil
}

func TestNewAssistant(t *testing.T) {

	assistantStore := NewAssistantStore()

	assistants, err := NewAssistant("test", "test", "test",false,false,false, assistantStore)
	defer RemoveAssistant("test", assistantStore)

	if err != nil {
		t.Errorf("NewAssistant failed: Didn't expect an error but got one, actual %s", err)
	}

	if len(*assistants) != 1 {
		t.Errorf("NewAssistant failed: Expected %d, actual %d", 1, len(*assistants))
	}

	if (*assistants)[0].Name != "test" {
		t.Errorf("NewAssistant failed: Expected %s, actual %s", "test", (*assistants)[0].Name)
	}

	allAssistants, error := ReadAssistants(assistantStore)

	if error != nil {
		t.Errorf("NewAssistant failed: Didn't expect an error but got one, actual %s", error)
	}

	if len(allAssistants) != 1 {
		t.Errorf("NewAssistant failed: Expected %d, actual %d", 1, len(allAssistants))
	}

}

func TestNewAssistantError(t *testing.T) {

	assistantStore := NewAssistantStore()

	_, err := NewAssistant("test", "test", "ErrorModel",false,false,false, assistantStore)

	if err == nil {
		t.Errorf("NewAssistant failed: Expected an error but didn't get one")
	}

}

func TestRemoveAssistant(t *testing.T) {

	assistantStore := NewAssistantStore()

	assistants, err := NewAssistant("test", "test", "test",false,false,false, assistantStore)

	if err != nil {
		t.Errorf("NewAssistant failed: Didn't expect an error but got one, actual %s", err)
	}
	if len(*assistants) != 1 {
		t.Errorf("RemoveAssistant failed: Expected %d, actual %d", 1, len(*assistants))
	}

	if (*assistants)[0].Name != "test" {
		t.Errorf("RemoveAssistant failed: Expected %s, actual %s", "test", (*assistants)[0].Name)
	}

	firstAssistantId := (*assistants)[0].ID

	updatedAssistants, error := RemoveAssistant(firstAssistantId, assistantStore)

	if error != nil {
		t.Errorf("NewAssistant failed: Didn't expect an error but got one, actual %s", error)
	}
	if len(updatedAssistants) != 0 {
		t.Errorf("RemoveAssistant failed: Expected %d, actual %d", 0, len(*assistants))
	}

	allAssistants, error := ReadAssistants(assistantStore)

	if error != nil {
		t.Errorf("NewAssistant failed: Didn't expect an error but got one, actual %s", error)
	}

	if len(allAssistants) != 0 {
		t.Errorf("RemoveAssistant failed: Expected %d, actual %d", 0, len(allAssistants))
	}
}

func TestUpdateAssistant(t *testing.T) {

	assistantStore := NewAssistantStore()

	assistants, err := NewAssistant("test", "test", "test",false,false,false, assistantStore)

	if err != nil {
		t.Errorf("NewAssistant failed: Didn't expect an error but got one, actual %s", err)
	}
	if len(*assistants) != 1 {
		t.Errorf("UpdateAssistant failed: Expected %d, actual %d", 1, len(*assistants))
	}

	if (*assistants)[0].Name != "test" {
		t.Errorf("UpdateAssistant failed: Expected %s, actual %s", "test", (*assistants)[0].Name)
	}

	firstAssistant := (*assistants)[0]

	updatedAssistants, error := UpdateAssistant(firstAssistant.ID, "test2", "test2", "test2",false,false,false, assistantStore)

	if error != nil {
		t.Errorf("UpdateAssistant failed: Didn't expect an error but got one, actual %s", error)
	}
	if len(updatedAssistants) != 1 {
		t.Errorf("UpdateAssistant failed: Expected %d, actual %d", 1, len(*assistants))
	}

	if (*assistants)[0].Name != "test2" {
		t.Errorf("UpdateAssistant failed: Expected %s, actual %s", "test2", (*assistants)[0].Name)
	}

	allAssistants, error := ReadAssistants(assistantStore)

	if error != nil {
		t.Errorf("UpdateAssistant failed: Didn't expect an error but got one, actual %s", error)
	}

	if len(allAssistants) != 1 {
		t.Errorf("UpdateAssistant failed: Expected %d, actual %d", 1, len(allAssistants))
	}

	if allAssistants[0].Name != "test2" {
		t.Errorf("UpdateAssistant failed: Expected %s, actual %s", "test2", allAssistants[0].Name)
	}

}

func TestFindAssistant(t *testing.T) {

	assistantStore := NewAssistantStore()

	NewAssistant("test", "test", "test",false,false,false, assistantStore)
	assistants, err := NewAssistant("test2", "test2", "test2",false,false,false, assistantStore)

	if err != nil {
		t.Errorf("NewAssistant failed: Didn't expect an error but got one, actual %s", err)
	}
	if len(*assistants) != 2 {
		t.Errorf("FindAssistant failed: Expected %d, actual %d", 2, len(*assistants))
	}

	if (*assistants)[0].Name != "test" {
		t.Errorf("FindAssistant failed: Expected %s, actual %s", "test", (*assistants)[0].Name)
	}

	firstAssistant := (*assistants)[0]
	secondAssistant := (*assistants)[1]

	foundAssistant, error := FindAssistant(firstAssistant.ID, assistantStore)

	if error != nil {
		t.Errorf("FindAssistant failed: Didn't expect an error but got one, actual %s", error)
	}

	if foundAssistant.Name != "test" {
		t.Errorf("FindAssistant failed: Expected %s, actual %s", "test", foundAssistant.Name)
	}

	foundAssistant, error = FindAssistant(secondAssistant.ID, assistantStore)

	if error != nil {
		t.Errorf("FindAssistant failed: Didn't expect an error but got one, actual %s", error)
	}

	if foundAssistant.Name != "test2" {
		t.Errorf("FindAssistant failed: Expected %s, actual %s", "test2", foundAssistant.Name)
	}

}

func TestFindAssistantError(t *testing.T) {

	assistantStore := NewAssistantStore()

	_, error := FindAssistant("test3", assistantStore)

	if error == nil {
		t.Errorf("FindAssistant failed: Expected an error but didn't get one")
	}

	_, error = NewAssistant("error", "test", "test",false,false,false, assistantStore)
	defer RemoveAssistant("test", assistantStore)

	if error != nil {
		t.Errorf("NewAssistant failed: Didn't expect an error but got one, actual %s", error)
	}

	_, error = FindAssistant("test3", assistantStore)

	if error == nil {
		t.Errorf("FindAssistant failed: Expected an error but didn't get one")
	}
}
