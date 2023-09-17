/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"assistants-cli/internal"

	"github.com/spf13/cobra"
)

var configKeyToReset string

// resetConfigCmd represents the resetConfig command
var resetConfigCmd = &cobra.Command{
	Use:   "resetConfig",
	Short: "Resets a config to its default value",
	Run: func(cmd *cobra.Command, args []string) {
		internal.ResetConfig(configKeyToReset)
	},
}

func init() {
	configCmd.AddCommand(resetConfigCmd)
	resetConfigCmd.Flags().StringVarP(&configKeyToReset, "key", "k", "", "The key of the config to reset")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
