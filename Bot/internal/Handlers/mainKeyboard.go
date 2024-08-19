package handler

import (
	"main/config"
	model "main/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MainMenuHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log := config.SetupLogger(config.Config("ENV"))

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = model.MainMenuKeyboard

	if _, err := bot.Send(msg); err != nil {
		log.Error("Error sending the message", "Err", err)
	}
}
