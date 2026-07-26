/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/khaingminhtun/api-inspector-cli/internal/config"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{

	Use: "get [url]",

	Short: "Send GET request",

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]

		timeout := config.Timeout()

		output := config.Output()

		fmt.Println("GET request to:", url)

		fmt.Printf(
			"URL: %s\nTimeout: %d\nOutput: %s\n",
			url,
			timeout,
			output,
		)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}
