package cmd

import (
	"github.com/khaingminhtun/api-inspector-cli/internal/client"
	"github.com/khaingminhtun/api-inspector-cli/internal/config"
	"github.com/khaingminhtun/api-inspector-cli/internal/formatter"
	"github.com/khaingminhtun/api-inspector-cli/internal/models"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post [url]",
	Short: "Send POST request",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		body, err := cmd.Flags().GetString("data")

		if err != nil {
			return err
		}

		request := models.Request{
			Method: "POST",
			URL:    args[0],
			Headers: config.MergeHeaders(
				appConfig.Headers,
				headers,
			),
			Body: body,
		}

		response, err := client.MakeRequest(
			request,
			appConfig.Timeout,
		)

		if err != nil {
			return err
		}

		return formatter.Print(
			appConfig.Output,
			response,
		)
	},
}

func init() {

	postCmd.Flags().StringP(
		"data",
		"d",
		"",
		"Request body",
	)

	rootCmd.AddCommand(postCmd)
}
