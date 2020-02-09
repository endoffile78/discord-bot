package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/endoffile78/bot/src/bot"
	"github.com/endoffile78/bot/src/config"
	"github.com/skwair/harmony"
)

var (
	configFile = ""
)

func init() {
	flag.StringVar(&configFile, "c", "settings.ini", "Settings file")
}

func main() {
	flag.Parse()

	fmt.Println("Starting bot...")

	err := config.ConfigLoad(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client, err := harmony.NewClient(config.ConfigGet("Discord", "token"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	bot := bot.NewBot(client)
	client.OnMessageCreate(bot.OnNewMessage)

	if err = client.Connect(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Bot started")

	defer client.Disconnect()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	config.ConfigSave(configFile)

	fmt.Println("Shutting down bot...")
}
