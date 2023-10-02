package cmd

import (
	"assistants-cli/internal"
	filesavers "assistants-cli/internal/fileSavers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var assistantId string
var assistantPrompt string
var assistantModel string
var assistantAllowSearch bool
var assistantAllowCommands bool
var assistantAllowFileReading bool

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the assistant",
	Long:  `Update the assistant file with new prompt and model`,
	Run: func(cmd *cobra.Command, args []string) {
		if assistantPrompt == "" && assistantModel == "" {
			fmt.Println("Nothing to update")
			os.Exit(1)
		}
		fileWriter := filesavers.NewAssistantFileStore(internal.ReadConfig(internal.AssistantFilePath))
		assistant, err := internal.FindAssistant(assistantId, fileWriter)
		if err != nil {
			fmt.Println("Error finding assistant:", err.Error())
			os.Exit(1)
		}
		if assistant == nil {
			fmt.Println("No assistant found with ID:", assistantId)
			os.Exit(1)
		}
		internal.UpdateAssistant(assistantId, assistant.Name, assistantPrompt, assistantModel, assistantAllowSearch, assistantAllowCommands, assistantAllowFileReading, fileWriter)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&assistantId, "assistantId", "a", "", "Name of the assistant")
	updateCmd.Flags().StringVarP(&assistantPrompt, "prompt", "p", "", "Prompt for the assistant")
	updateCmd.Flags().StringVarP(&assistantModel, "model", "m", "", "Default Model to use with the assistant")
	updateCmd.Flags().BoolVarP(&assistantAllowSearch, "allow-search", "s", false, "Allow the assistant to search the web")
	updateCmd.Flags().BoolVarP(&assistantAllowCommands, "allow-commands", "c", false, "Allow the assistant to run commands")
	updateCmd.Flags().BoolVarP(&assistantAllowFileReading, "allow-file-reading", "f", false, "Allow the assistant to read files")
	updateCmd.MarkFlagRequired("assistantId")

}
