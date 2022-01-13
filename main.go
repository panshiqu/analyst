package main

import (
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/natefinch/lumberjack"
	"github.com/panshiqu/analyst/define"
	"github.com/panshiqu/analyst/handler"
	"github.com/panshiqu/analyst/utils"
)

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename: "./log/analyst.log",
		MaxSize:  50,
		MaxAge:   30,
	})

	if err := utils.ReadJSON("config.json", &define.C); err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(define.C.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Println("Authorized on account", bot.Self.UserName)

	updates := bot.ListenForWebhook(fmt.Sprintf("/%s", bot.Token))
	go http.ListenAndServe(":8443", nil)

	for u := range updates {
		if u.Message == nil {
			continue
		}

		log.Printf("Recv [%s] %s\n", u.Message.From.UserName, u.Message.Text)

		text, err := handler.Handle(u.Message.From.UserName, u.Message.Text)
		if err != nil {
			text = err.Error()
		}

		log.Printf("Send [%s] %s\n", u.Message.From.UserName, text)

		msg := tgbotapi.NewMessage(u.Message.Chat.ID, text)

		if _, err := bot.Send(msg); err != nil {
			log.Println("Send", err)
		}
	}
}
