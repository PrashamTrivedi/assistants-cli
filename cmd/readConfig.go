/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"assistants-cli/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var configKey string

// readConfigCmd represents the readConfig command
var readConfigCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads the current configuration",

	Run: func(cmd *cobra.Command, args []string) {
		configs := internal.ReadAllConfig()
		if configKey != "" {
			fmt.Println(configs[configKey])
			return
		}
		for key, value := range configs {
			fmt.Printf("%s: %s\n", key, value)
		}

	},
}

func init() {
	configCmd.AddCommand(readConfigCmd)
	readConfigCmd.Flags().StringVarP(&configKey, "key", "k", "", "The key of the config to read")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
