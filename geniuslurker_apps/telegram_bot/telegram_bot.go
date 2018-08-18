package main

import (
	"flag"

	"github.com/AlexanderYAPPO/geniuslurker/geniuslurker"
	"github.com/bot-api/telegram"
	"github.com/bot-api/telegram/telebot"
	"golang.org/x/net/context"
)

func main() {
	geniuslurker.InitLoggers()
	geniuslurker.InitSettings()

	token := flag.String("token", "", "telegram bot token")
	debug := flag.Bool("debug", false, "show debug information")
	flag.Parse()

	if *token == "" {
		geniuslurker.ErrorLogger.Fatalln("token flag is required")
	}

	api := telegram.New(*token)
	api.Debug(*debug)
	bot := telebot.NewWithAPI(api)
	bot.Use(telebot.Recover()) // recover if handler panic

	netCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Use command middleware, that helps to work with commands
	bot.Use(telebot.Commands(map[string]telebot.Commander{
		"search": telebot.CommandFunc(geniuslurker.SearchCommand),
		"lyrics": telebot.CommandFunc(geniuslurker.GetLyricsCommand),
	}))

	err := bot.Serve(netCtx)
	if err != nil {
		geniuslurker.ErrorLogger.Fatalln(err)
	}
	geniuslurker.InfoLogger.Println("Telegram bot started")
}
