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
}

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.Flags().StringVarP(&key, "key", "k", "", "configuration key (required)")
	rootCmd.MarkFlagRequired("key")
	viper.BindPFlag("openai_key", rootCmd.Flags().Lookup("key"))
}
