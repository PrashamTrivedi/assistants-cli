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

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the assistant",
	Long:  `Update the assistant file with new prompt and model`,
	Run: func(cmd *cobra.Command, args []string) {
		if assistantPrompt == "" && assistantModel == "" {
			fmt.Println("Nothing to update")
			os.Exit(1)
		}
		fileWriter := filesavers.NewAssistantFileStore(internal.AssistantFilePath)
		assistant, err := internal.FindAssistant(assistantId, fileWriter)
		if err != nil {
			fmt.Println("Error finding assistant:", err.Error())
			os.Exit(1)
		}
		if assistant == nil {
			fmt.Println("No assistant found with ID:", assistantId)
			os.Exit(1)
		}
		internal.UpdateAssistant(assistantId, assistant.Name, assistantPrompt, assistantModel, fileWriter)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&assistantId, "assistantId", "a", "", "Name of the assistant")
	updateCmd.Flags().StringVarP(&assistantPrompt, "prompt", "p", "", "Prompt for the assistant")
	updateCmd.Flags().StringVarP(&assistantModel, "model", "m", "", "Default Model to use with the assistant")
	updateCmd.MarkFlagRequired("assistantId")

}
