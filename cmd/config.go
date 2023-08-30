package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var key string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set OpenAI Key",
	Long:  `Set OpenAI key to use for API calls.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigFile("config.json")
		viper.Set("openai_key", key)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&key, "key", "k", "", "configuration key (required)")
	configCmd.MarkFlagRequired("key")

}
