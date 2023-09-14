package cmd

import (
	"assistants-cli/internal"
	filesavers "assistants-cli/internal/fileSavers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var assistantIdForChat string
var chatId string
var continueChat bool
var message string

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with one of the assistants",
	Long: `This command will either start or continue a chat with one of the assistants. 
	You can start a new chat by providing the assistant ID.
	You can continue a chat by providing the chat ID or passing continue flag which will continue with latest chat ids.`,

	Run: func(cmd *cobra.Command, args []string) {

		fileWriter := filesavers.NewAssistantFileStore(internal.ReadConfig(internal.AssistantFilePath))
		chatFileWriter := filesavers.NewChatFileStore(internal.ReadConfig(internal.ChatFilePath))
		if continueChat {
			chatId = internal.ReadConfig("latest_chat_id")
		}
		chatIdToStore := ""

		if message == "" {
			if continueChat || chatId != "" {
				fmt.Println("Message is required")
				os.Exit(1)
			}
			message = "Hello There!"
		}
		openaiKey := viper.GetString("openai_key")
		if chatId == "" {
			assistant, error := internal.FindAssistant(assistantIdForChat, fileWriter)
			if error != nil {
				fmt.Println("No assistant found with ID:", assistantIdForChat)
				os.Exit(1)
			}
			chat, error := internal.NewChat(internal.ApiKey(openaiKey), *assistant)
			if error != nil {
				fmt.Println("Error creating chat:", error.Error())
				os.Exit(1)
			}
			chatIdFromStorage, err := chat.Start(message, chatFileWriter)
			if err != nil {
				fmt.Println("Error starting chat:", err.Error())
				os.Exit(1)
			}
			chatIdToStore = chatIdFromStorage
		} else {
			chat, error := internal.GetChat(chatId, internal.ApiKey(openaiKey), chatFileWriter, fileWriter)
			if error != nil {
				fmt.Println("Error getting chat:", error.Error())
				os.Exit(1)
			}
			err := chat.Continue(message, chatFileWriter, fileWriter)
			if err != nil {
				fmt.Println("Error continuing chat:", err.Error())
				os.Exit(1)
			}
			chatIdToStore = chatId

		}
		internal.WriteConfig(map[string]string{
			"latest_chat_id": chatIdToStore,
		})
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().StringVarP(&assistantIdForChat, "assistantId", "a", "", "Name of the assistant")
	chatCmd.Flags().StringVarP(&chatId, "chatId", "c", "", "Chat ID to continue")
	chatCmd.Flags().BoolVarP(&continueChat, "continue", "", false, "Continue with latest chat ID")
	chatCmd.Flags().StringVarP(&message, "message", "m", "", "Message to send to the assistant")

	chatCmd.MarkFlagsMutuallyExclusive("chatId", "assistantId")
	chatCmd.MarkFlagsMutuallyExclusive("chatId", "continue")
	chatCmd.MarkFlagsMutuallyExclusive("assistantId", "continue")

}
