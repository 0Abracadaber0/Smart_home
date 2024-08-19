package handler

import (
	"main/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessagesHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log := config.SetupLogger(config.Config("ENV"))

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Text {
	case "Температура":
		msg.Text = "Это температура!"
	case "Розетки":
		msg.Text = "Это розетки!"
	case "Безопасность":
		msg.Text = "Это безопасность!"
	case "Настройки":
		msg.Text = "Это настройки!"
	default:
		msg.Text = "Не знаю такого("
	}
	if _, err := bot.Send(msg); err != nil {
		log.Error("Error sending the message", "Err", err)
	}
}
