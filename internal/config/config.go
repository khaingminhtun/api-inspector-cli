package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Load() error {

	home, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	viper.AddConfigPath(home)

	viper.SetConfigName(".apispy")

	viper.SetConfigType("yaml")

	//Defaults

	viper.SetDefault(
		"timeout",
		30,
	)

	viper.SetDefault(
		"output",
		"json",
	)

	viper.SetDefault(
		"headers",
		map[string]string{
			"User-Agent": "apispy/1.0",
		},
	)

	// Environment

	viper.SetEnvPrefix("APISPY")

	viper.AutomaticEnv()

	// Read config

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println(
			"Using default configuration",
		)
		return nil
	}

	return nil
}

func Timeout() int {
	return viper.GetInt("timeout")
}

func Output() string {
	return viper.GetString("output")
}
