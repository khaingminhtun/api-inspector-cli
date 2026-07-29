package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct{

    Timeout int

	Output string

	Headers map[string]string
}

func Load() (*Config,error) {

	home, err := os.UserHomeDir()

	if err != nil {
		return nil,err
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
		fmt.Println(
			"Using default configuration",
		)
		
	}

	 cfg := &Config{

		Timeout: viper.GetInt("timeout"),

		Output: viper.GetString("output"),

		Headers: viper.GetStringMapString("headers"),
	 }

	 return cfg, nil
}


