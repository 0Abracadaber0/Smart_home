package handler

import (
	"main/config"
	model "main/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var currentMenu = "main"

func MessagesHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log := config.SetupLogger(config.Config("ENV"))

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	switch currentMenu {
	case "main":
		switch update.Message.Text {
		case "Температура":
			msg.Text = "Это температура!"
		case "Розетки":
			msg.Text = "Это розетки!"
		case "Безопасность":
			msg.Text = "Это безопасность!"
		case "Настройки":
			currentMenu = "settings"
			KeyboardHandler(bot, update, model.SettingsKeyboard)
		default:
			msg.Text = "Не знаю такого("
		}
	case "settings":
		switch update.Message.Text {
		case "Перезапуск шлюза":
			msg.Text = "Это ребут шлюза!"
		case "Помощь":
			msg.Text = "Это помощь!"
		case "Админ":
			msg.Text = "Это админ панель!"
		case "Вернуться в главное меню":
			currentMenu = "main"
			KeyboardHandler(bot, update, model.MainMenuKeyboard)
		default:
			msg.Text = "Не знаю такого("
		}
	}

	if _, err := bot.Send(msg); err != nil {
		log.Error("Error sending the message", "Err", err)
	}
}
