package cmd

import (
	"fmt"
	"os"

	"github.com/khaingminhtun/api-inspector-cli/internal/config"
	"github.com/spf13/cobra"
)

var appConfig *config.Config

var headers []string

var rootCmd = &cobra.Command{
	Use:   "apispy",
	Short: "API Inspection CLI",
	Long: `
API Spy is a CLI tool for inspecting HTTP APIs.
	`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return applyFlags(cmd)
	},
}

func SetConfig(cfg *config.Config) {
	appConfig = cfg
}

func Execute() {

	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func applyFlags(cmd *cobra.Command) error {

	timeout, err := cmd.Flags().GetInt("timeout")

	if err != nil {
		return err
	}

	if timeout != 0 {
		appConfig.Timeout = timeout
	}

	output, err := cmd.Flags().GetString("output")

	if err != nil {
		return err
	}

	if output != "" {
		appConfig.Output = output
	}

	return nil
}

func init() {

	rootCmd.PersistentFlags().
		Int(
			"timeout",
			0,
			"request timeout",
		)

	rootCmd.PersistentFlags().
		String(
			"output",
			"",
			"output format",
		)

	rootCmd.PersistentFlags().
		StringSliceVarP(
			&headers,
			"header",
			"H",
			[]string{},
			"HTTP headers",
		)

}
