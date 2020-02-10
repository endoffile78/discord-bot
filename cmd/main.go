package main

import (
	"flag"
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/endoffile78/bot/bot"
	"github.com/endoffile78/bot/config"
)

var (
	configFile = ""
)

func init() {
	flag.StringVar(&configFile, "c", "settings.ini", "Settings file")
}

func main() {
	flag.Parse()

	err := config.ConfigLoad(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := disgord.New(disgord.Config{
		BotToken: config.ConfigGet("Discord", "token"),
	})

	bot := bot.NewBot(client)
	bot.Run()

	fmt.Println("Saving to", configFile)
	config.ConfigSave(configFile)
}
