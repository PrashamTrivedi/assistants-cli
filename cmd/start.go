package cmd

import (
	"assistants-cli/internal"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var assistantNameForChat string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the chat assistant",
	Long:  `This command starts the chat assistant and prompts the user for input.`,
	Run: func(cmd *cobra.Command, args []string) {

		fileWriter := internal.NewFileStore("assistants.json")
		assistant, error := internal.FindAssistant(assistantNameForChat, fileWriter)
		if error != nil {
			fmt.Println("No assistant found with name:", assistantNameForChat)
			os.Exit(1)
		}
		openaiKey := viper.GetString("openai_key")
		chat, error := internal.NewChat(internal.ApiKey(openaiKey), *assistant)
		if error != nil {
			fmt.Println("Error creating chat:", error.Error())
			os.Exit(1)
		}
		chat.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&assistantNameForChat, "assistant", "a", "", "Name of the assistant")
	startCmd.MarkFlagRequired("assistant")
}
