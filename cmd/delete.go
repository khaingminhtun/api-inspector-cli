package cmd

import (
	"fmt"

	"github.com/khaingminhtun/api-inspector-cli/internal/storage"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{

	Use: "delete [name]",

	Short: "Delete saved request",

	Args: cobra.ExactArgs(1),

	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {

		name := args[0]

		err := storage.Delete(name)

		if err != nil {
			return err
		}

		fmt.Println(
			"Deleted:",
			name,
		)

		return nil
	},
}

func init() {

	rootCmd.AddCommand(deleteCmd)

}
