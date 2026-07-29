/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/khaingminhtun/api-inspector-cli/cmd"
	"github.com/khaingminhtun/api-inspector-cli/internal/config"
)

func main() {

	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	cmd.SetConfig(cfg)
	cmd.Execute()
}
