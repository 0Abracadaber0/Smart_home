package handler

import (
	"main/config"
	model "main/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CommandsHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, botState *model.BotState) {
	log := config.SetupLogger(config.Config("ENV"))

	switch update.Message.Command() {
	case "menu":
		botState.CurrentMenu = "main"
		KeyboardHandler(bot, update, model.MainMenuKeyboard)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не знаю такого(")
		if _, err := bot.Send(msg); err != nil {
			log.Error("Error sending the message", "Err", err)
		}
	}
}
