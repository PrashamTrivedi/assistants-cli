/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"assistants-cli/internal"
	filesavers "assistants-cli/internal/fileSavers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Assistants",
	Long:  `List all your assistants`,
	Run: func(cmd *cobra.Command, args []string) {

		fileWriter := filesavers.NewAssistantFileStore(internal.ReadConfig(internal.AssistantFilePath))
		assistants, err := internal.ReadAssistants(fileWriter)
		if err != nil {
			fmt.Println("Error reading assistants:", err.Error())
			os.Exit(1)
		}
		if len(assistants) == 0 {
			fmt.Println("You have no assistants")
			os.Exit(0)
		}
		for i, assistant := range assistants {
			fmt.Printf("%d. %s\n", i+1, assistant.ID)
			fmt.Printf("\t Name: %s\n", assistant.Name)
			fmt.Printf("\t Prompt: %s\n", assistant.Prompt)
			fmt.Printf("\t Default Model: %s\n", assistant.DefaultModel)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
