package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	AssistantFilePath = "assistant_file_path"
	ChatFilePath      = "chat_file_path"
	OpenaiKey         = "openai_key"
)

func ReadConfig(configKey string) string {
	configPath := createConfigPath()
	viper.SetConfigFile(configPath)
	viper.ReadInConfig()
	return viper.GetString(configKey)
}

func ReadAllConfig() map[string]interface{} {
	configPath := createConfigPath()
	viper.SetConfigFile(configPath)
	viper.ReadInConfig()
	return viper.AllSettings()
}
func ResetConfig(key string) {
	configPath := createConfigPath()
	viper.SetConfigFile(configPath)
	if key == "" {
		viper.WriteConfig()
		return
	}
	viper.ReadInConfig()
	viper.Set(key, "")
	viper.WriteConfig()
}
func createConfigPath() string {
	if secretPath, exists := os.LookupEnv("SECRET_CONFIG_PATH"); exists {
		configFilePath := secretPath
		fmt.Println("Warning: You are using a secret config path meant for testing. Proceed with caution.")
		return configFilePath
	}
	home, _ := os.UserHomeDir()
	configPath := filepath.Join(home, ".assistants", "config.json")
	// Create ~/.assistants/config.json if it doesn't exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.MkdirAll(filepath.Join(home, ".assistants"), os.ModePerm)
		os.Create(configPath)
	}
	return configPath
}

func WriteConfig(keyVauleMap map[string]string) {
	configPath := createConfigPath()

	viper.SetConfigFile(configPath)

	for key, value := range keyVauleMap {
		viper.Set(key, value)
	}
	if err := viper.WriteConfig(); err != nil {
		panic(err)
	}
}
