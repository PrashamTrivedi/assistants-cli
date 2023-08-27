package cmd

import (
	"assistants-cli/internal"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var assistantName string
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
		fileWriter := internal.NewFileStore("assistants.json")

		internal.UpdateAssistant(assistantName, assistantPrompt, assistantModel, fileWriter)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&assistantName, "name", "n", "", "Name of the assistant")
	updateCmd.Flags().StringVarP(&assistantPrompt, "prompt", "p", "", "Prompt for the assistant")
	updateCmd.Flags().StringVarP(&assistantModel, "model", "m", "", "Default Model to use with the assistant")
	updateCmd.MarkFlagRequired("name")

}
