package cmd

import (
	"assistants-cli/internal"

	"github.com/spf13/cobra"
)

var key string
var assistantFilePath string
var chatFilePath string
var serpApiKey string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set OpenAI Key and file paths to store assistants and chats",
	Long:  `Set OpenAI key to use for API calls. File paths are optional and will default to ~/.assistants/assistants.json and ~/.assistants/chats.json (In windows this will be \%USERPROFILE%\.assistants\assistants.json and \%USERPROFILE%\.assistants\chats.json)`,
	Run: func(cmd *cobra.Command, args []string) {
		configToSet := map[string]string{}
		configToSet[internal.OpenaiKey] = key
		if assistantFilePath != "" {
			configToSet[internal.AssistantFilePath] = assistantFilePath
		}
		if chatFilePath != "" {
			configToSet[internal.ChatFilePath] = chatFilePath
		}
		configToSet[internal.SerpApiKey] = serpApiKey
		internal.WriteConfig(configToSet)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&key, "key", "k", "", "configuration key (required)")
	configCmd.Flags().StringVarP(&assistantFilePath, "assistantFilePath", "a", "", "path to assistant file")
	configCmd.Flags().StringVarP(&chatFilePath, "chatFilePath", "c", "", "path to chat file")
	configCmd.Flags().StringVarP(&serpApiKey, "serpApiKey", "s", "", "serp api key")
	configCmd.MarkFlagRequired("key")

}
