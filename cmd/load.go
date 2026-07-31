package cmd

import (
	"github.com/khaingminhtun/api-inspector-cli/internal/client"
	"github.com/khaingminhtun/api-inspector-cli/internal/formatter"
	"github.com/khaingminhtun/api-inspector-cli/internal/storage"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{

	Use: "load [name]",

	Short: "Load saved request",

	Args: cobra.ExactArgs(1),

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {

		name := args[0]

		// Load request from storage

		request, err := storage.Load(name)

		if err != nil {
			return err
		}

		// Execute request

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

	rootCmd.AddCommand(loadCmd)

}
