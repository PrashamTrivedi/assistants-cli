/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"assistants-cli/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var configKeyToReset string

// resetConfigCmd represents the resetConfig command
var resetConfigCmd = &cobra.Command{
	Use:   "resetConfig",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("resetConfig called")
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
