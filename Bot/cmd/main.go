package main

import (
	"fmt"
	"main/config"
	"main/internal/connect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	log := config.SetupLogger(config.Config("ENV"))
	log.Info("starting server", "env", config.Config("ENV"))

	bot := connect.ConnectBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Info(fmt.Sprintf("[%s] %s", update.Message.From.UserName, update.Message.Text))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

}
