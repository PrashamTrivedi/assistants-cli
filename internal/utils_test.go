package internal

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func Test_createConfigPathDefault(t *testing.T) {
	configPath := createConfigPath()
	if configPath != "/root/.assistants/config.json" {
		t.Error("createConfigPathDefault failed")
	}
}

func Test_createConfigPathSecret(t *testing.T) {
	secretPath := "/workspaces/assistants-cli/test/.assistants/config.json"
	os.Setenv("SECRET_CONFIG_PATH", secretPath)
	defer os.Unsetenv("SECRET_CONFIG_PATH")
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
}

func Test_WriteConfigForSetupAndTearDown(t *testing.T) {
	secretPath := "/workspaces/assistants-cli/test/.assistants/config.json"
	os.Setenv("SECRET_CONFIG_PATH", secretPath)
	defer os.Unsetenv("SECRET_CONFIG_PATH")
	defer ResetConfig("test")
	defer viper.Reset()
	WriteDefaultValues()
	testKeyVal := make(map[string]string)
	testKeyVal["test"] = "test"
	WriteConfig(testKeyVal)
	config := ReadConfig("test")
	if config != "test" {
		t.Errorf("WriteConfig failed: Expected %s, actual %s", "test", config)
	}
}

func WriteDefaultValues() {
	ResetConfig("")
	keyValueMap := make(map[string]string)
	keyValueMap[AssistantFilePath] = "X"
	keyValueMap[ChatFilePath] = "Y"
	keyValueMap[OpenaiKey] = "Z"
	WriteConfig(keyValueMap)
}

func Test_ReadAllConfig(t *testing.T) {
	secretPath := "/workspaces/assistants-cli/test/.assistants/config.json"
	os.Setenv("SECRET_CONFIG_PATH", secretPath)
	defer os.Unsetenv("SECRET_CONFIG_PATH")
	WriteDefaultValues()
	configMap := ReadAllConfig()
	for key, value := range configMap {
		if value == "" {
			configMap[key] = nil
		}
	}
	if len(configMap) != 3 {
		t.Logf("ReadAllConfig failed, Expected length %d, actual length %d", 3, len(configMap))
	}
}

func Test_ReadConfigUnknown(t *testing.T) {
	secretPath := "/workspaces/assistants-cli/test/.assistants/config.json"
	os.Setenv("SECRET_CONFIG_PATH", secretPath)
	defer os.Unsetenv("SECRET_CONFIG_PATH")
	WriteDefaultValues()
	config := ReadConfig("unknown")
	if config != "" {
		t.Error("ReadConfigUnknown failed: It should be blank")
	}
}

func Test_WriteConfig(t *testing.T) {
	secretPath := "/workspaces/assistants-cli/test/.assistants/config.json"
	os.Setenv("SECRET_CONFIG_PATH", secretPath)
	defer os.Unsetenv("SECRET_CONFIG_PATH")
	defer ResetConfig("test")
	WriteDefaultValues()
	keyValueMap := make(map[string]string)
	keyValueMap["test"] = "test"
	WriteConfig(keyValueMap)
	config := ReadConfig("test")
	if config != "test" {
		t.Errorf("WriteConfig failed: Expected %s, actual %s", "test", config)
	}
}

func Test_ResetConfig(t *testing.T) {
	secretPath := "/workspaces/assistants-cli/test/.assistants/config.json"
	os.Setenv("SECRET_CONFIG_PATH", secretPath)
	defer os.Unsetenv("SECRET_CONFIG_PATH")
	WriteDefaultValues()
	keyValueMap := make(map[string]string)
	keyValueMap["test"] = "test"
	keyValueMap["test2"] = "test2"
	WriteConfig(keyValueMap)
	ResetConfig("test")
	config := ReadConfig("test")
	if config != "" {
		t.Errorf("ResetConfig failed: Expected %s, actual %s", "", config)
	}
	// Test that configs should not have test key
	updatedConfigs := ReadAllConfig()
	if value, ok := updatedConfigs["test"]; ok && value != "" {
		t.Errorf("ResetConfig failed: Expected %s, actual %s", "", config)
	}

	if value, ok := updatedConfigs["test2"]; !ok || value != "test2" {
		t.Errorf("ResetConfig failed for other keys: Expected %s, actual %s", "test2", "")
	}

	ResetConfig("test2")

}
