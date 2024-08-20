package handler

import (
	"main/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func KeyboardHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, keyboard tgbotapi.ReplyKeyboardMarkup) {
	log := config.SetupLogger(config.Config("ENV"))

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню:")
	msg.ReplyMarkup = keyboard

	if _, err := bot.Send(msg); err != nil {
		log.Error("Error sending the message", "Err", err)
	}
}
