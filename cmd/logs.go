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

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Shows your chat history",
	Long:  `Shows your chat history, including ChatID and assistant ID`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logs called")
		chatFileWriter := filesavers.NewChatFileStore("chats.json")
		chatData, err := internal.ListChat(chatFileWriter)
		if err != nil {
			fmt.Println("Error reading chat:", err.Error())
			os.Exit(1)
		}

		if len(chatData) == 0 {
			fmt.Println("You have no chat history")
			os.Exit(0)
		}

		for i, chat := range chatData {
			fmt.Printf("%d. %s\n", i+1, chat.ID)
			fmt.Printf("\t Assistant ID: %s\n", chat.AssistantId)
			fmt.Printf("\t Messages: %s\n", chat.Messages)
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
