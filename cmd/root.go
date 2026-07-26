/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/khaingminhtun/api-inspector-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	timeout int

	output string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apispy",
	Short: "API Inspection CLI",
	Long: `
API Spy is a CLI tool for inspecting HTTP APIs.
	`,
}

func Execute() {

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().
		IntVar(
			&timeout,
			"timeout",
			30,
			"request timeout",
		)

	viper.BindPFlag(
		"timeout",
		rootCmd.PersistentFlags().Lookup("timeout"),
	)

	rootCmd.PersistentFlags().
		StringVar(
			&output,
			"output",
			"json",
			"output format",
		)

	viper.BindPFlag(
		"output",
		rootCmd.PersistentFlags().Lookup("output"),
	)

}
