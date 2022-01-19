package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/natefinch/lumberjack"
	"github.com/panshiqu/analyst/cache"
	"github.com/panshiqu/analyst/define"
	"github.com/panshiqu/analyst/handler"
	"github.com/panshiqu/analyst/utils"
	"github.com/robfig/cron/v3"
)

func handleSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	s := <-c
	log.Println("Got signal:", s)

	if err := cache.SaveAlerts(); err != nil {
		log.Println("SaveAlerts", err)
	}

	if err := cache.SavePrices(); err != nil {
		log.Println("SavePrices", err)
	}

	os.Exit(0)
}

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename: "./log/analyst.log",
		MaxSize:  50,
		MaxAge:   30,
	})

	if err := cache.LoadAlerts(); err != nil {
		log.Fatal(err)
	}

	if err := cache.LoadPrices(); err != nil {
		log.Fatal(err)
	}

	if err := utils.ReadJSON("config.json", &define.C); err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(define.C.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	cache.Bot = bot

	log.Println("Authorized on account", bot.Self.UserName)

	updates := bot.ListenForWebhook(fmt.Sprintf("/%s", bot.Token))

	http.HandleFunc("/prices", cache.ServeHTTP)

	go http.ListenAndServe(":8443", nil)

	go handleSignal()

	c := cron.New()
	c.AddFunc("* * * * *", handler.PerMinute)
	c.Start()

	for u := range updates {
		if u.Message == nil {
			continue
		}

		log.Printf("Recv [%d] [%s] %s\n", u.Message.Chat.ID, u.Message.From.UserName, u.Message.Text)

		text, err := handler.Handle(u.Message.Chat.ID, u.Message.From.UserName, u.Message.Text)
		if err != nil {
			text = err.Error()
		}

		log.Printf("Send [%d] [%s] %s\n", u.Message.Chat.ID, u.Message.From.UserName, text)

		msg := tgbotapi.NewMessage(u.Message.Chat.ID, text)

		if _, err := bot.Send(msg); err != nil {
			log.Println("Send", err)
		}
	}
}
