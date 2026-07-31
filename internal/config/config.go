package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Timeout int

	Output string

	Headers map[string]string
}

func Load() (*Config, error) {

	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err
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

	if err := viper.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Using default configuration")
		} else {
			return nil, err
		}
	}

	cfg := &Config{

		Timeout: viper.GetInt("timeout"),

		Output: viper.GetString("output"),

		Headers: viper.GetStringMapString("headers"),
	}

	return cfg, nil
}

func MergeHeaders(
	defaults map[string]string,
	values []string,
) map[string]string {

	result := make(map[string]string)

	for key, value := range defaults {
		result[key] = value
	}

	for _, header := range values {

		parts := strings.SplitN(
			header,
			":",
			2,
		)

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])

		value := strings.TrimSpace(parts[1])

		result[key] = value
	}

	return result
}
