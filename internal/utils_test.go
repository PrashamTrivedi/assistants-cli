package internal

import (
	"fmt"
	"os"
	"testing"
)

func Test_createConfigPathDefault(t *testing.T) {
	configPath := createConfigPath()
	if configPath != "/root/.assistants/config.json" {
		t.Error("createConfigPathDefault failed")
	}
}

func Test_createConfigPathSecret(t *testing.T) {
	secretPath := "~/testing/.assistants/secret.json"
	os.Setenv("SECRET_CONFIG_PATH", secretPath)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current Directory:", dir)
	configPath := createConfigPath()
	if configPath != secretPath {
		t.Error("createConfigPathSecret failed")
	}
	os.Unsetenv("SECRET_CONFIG_PATH")
}

func Test_ReadAllConfig(t *testing.T) {
	secretPath := "~/testing/.assistants/secret.json"
	os.Setenv("SECRET_CONFIG_PATH", secretPath)
	ResetConfig("")
	configMap := ReadAllConfig()
	if len(configMap) != 3 {
		t.Error("ReadAllConfig failed")
	}
}
