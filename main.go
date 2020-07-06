package main

import (
	"anchorfm-bot/infra"
	"anchorfm-bot/telegram"
	"os"
)

func main() {
	var configPath string

	if len(os.Args) == 2 {
		configPath = os.Args[1]
	} else {
		configPath = os.Getenv("ConfigPath")
	}

	config, err := infra.Configuration{}.Load(configPath)
	if err != nil {
		panic(err)
	}

	telegram.Run(config)
}