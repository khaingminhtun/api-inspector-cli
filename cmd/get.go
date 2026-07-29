/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/khaingminhtun/api-inspector-cli/internal/client"
	"github.com/khaingminhtun/api-inspector-cli/internal/formatter"
	"github.com/khaingminhtun/api-inspector-cli/internal/models"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{

	Use: "get [url]",

	Short: "Send GET request",

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		request := models.Request{
			Method:  "GET",
			URL:     args[0],
			Headers: appConfig.Headers,
		}

		response, err := client.MakeRequest(
			request,
			appConfig.Timeout,
		)

		if err != nil {
			return err
		}

		err = formatter.Print(
			appConfig.Output,
			response,
		)

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}
