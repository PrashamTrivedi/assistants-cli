package cmd

import (
	"assistants-cli/internal"
	filesavers "assistants-cli/internal/fileSavers"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var name string
var prompt string
var model string
var allowSearch bool
var allowCommands bool
var allowFileReading bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new assistant",
	Long:  `Create a new assistant with the specified prompt, model and name.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileWriter := filesavers.NewAssistantFileStore(internal.ReadConfig(internal.AssistantFilePath))
		if model == "" {
			model = openai.GPT3Dot5Turbo16K
		}
		_, err := internal.NewAssistant(name, prompt, model, allowSearch, allowCommands, allowFileReading, fileWriter)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the assistant")
	createCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Prompt for the assistant")
	createCmd.Flags().StringVarP(&model, "model", "m", "", "Default Model to use with the assistant")
	createCmd.Flags().BoolVarP(&allowSearch, "allow-search", "s", false, "Allow the assistant to search the web")
	createCmd.Flags().BoolVarP(&allowCommands, "allow-commands", "c", false, "Allow the assistant to run commands")
	createCmd.Flags().BoolVarP(&allowFileReading, "allow-file-reading", "f", false, "Allow the assistant to read files")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("prompt")
}
