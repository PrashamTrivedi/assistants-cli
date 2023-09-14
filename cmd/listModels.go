/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"assistants-cli/internal"
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

// listModelsCmd represents the listModels command
var listModelsCmd = &cobra.Command{
	Use:   "listModels",
	Short: "List Available models",
	Long:  `List available models for the OpenAI API. This is useful to create assistants with different models.`,
	Run: func(cmd *cobra.Command, args []string) {

		openaiKey := internal.ReadConfig(internal.OpenaiKey)
		if openaiKey == "" {
			fmt.Println("OpenAI key not found. Please set it using the config command.")
			os.Exit(1)
		}
		client := openai.NewClient(openaiKey)
		fmt.Println("Listing models...")
		models, err := client.ListModels(context.Background())

		// Sort models such a way that models containing gpt are listed first
		sort.Slice(models.Models, func(i, j int) bool {

			return models.Models[j].CreatedAt < models.Models[i].CreatedAt
		})
		if err != nil {
			fmt.Println("Error listing models:", err.Error())
		}
		for i, model := range models.Models {
			fmt.Printf("%d, %s\n", i, model.ID)
			// fmt.Printf("%+v\n", model)
		}

	},
}

func init() {
	rootCmd.AddCommand(listModelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listModelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listModelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
