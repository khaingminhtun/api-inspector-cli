package cmd

import (
	"fmt"

	"github.com/khaingminhtun/api-inspector-cli/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{

	Use: "list",

	Short: "List saved requests",

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {

		requests, err := storage.List()

		if err != nil {
			return err
		}

		fmt.Println("Saved Requests:")

		fmt.Println()

		for _, name := range requests {

			fmt.Println(name)

		}

		return nil
	},
}

func init() {

	rootCmd.AddCommand(listCmd)

}
