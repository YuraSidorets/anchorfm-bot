package main

import (
	"anchorfm-bot/infra"
	"anchorfm-bot/telegram"
	"os"
)

func main() {
	var configPath string
	var env bool
	if len(os.Args) == 2 {
		configPath = os.Args[1]
		env = false
	} else {
		configPath = os.Getenv("ConfigPath")
		env = true
	}

	config, err := infra.Configuration{}.Load(configPath, env)
	if err != nil {
		panic(err)
	}

	telegram.Run(config)
}