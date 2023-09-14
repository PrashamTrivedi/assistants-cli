package cmd

import (
	"assistants-cli/internal"
	filesavers "assistants-cli/internal/fileSavers"

	"github.com/spf13/cobra"
)

var assistantNameToRemove string

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a resource",
	Long:  `Remove a resource with the specified name.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileWriter := filesavers.NewAssistantFileStore(internal.ReadConfig(internal.AssistantFilePath))
		internal.RemoveAssistant(assistantNameToRemove, fileWriter)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVarP(&assistantNameToRemove, "name", "n", "", "Name of the assistant")
	removeCmd.MarkFlagRequired("name")
}
