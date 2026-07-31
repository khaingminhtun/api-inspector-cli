package cmd

import (
	"fmt"

	"github.com/khaingminhtun/api-inspector-cli/internal/models"
	"github.com/khaingminhtun/api-inspector-cli/internal/storage"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{

	Use: "save [name] [url]",

	Short: "Save request",

	Args: cobra.ExactArgs(2),

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {

		request := models.Request{

			Method: "GET",

			URL: args[1],

			Headers: appConfig.Headers,
		}

		err := storage.Save(
			args[0],
			request,
		)

		if err != nil {
			return err
		}

		fmt.Println(
			"Saved:",
			args[0],
		)

		return nil
	},
}

func init() {

	rootCmd.AddCommand(saveCmd)

}
