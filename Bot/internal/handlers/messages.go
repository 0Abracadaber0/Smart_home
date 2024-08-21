package handler

import (
	"main/config"
	model "main/internal/models"
	request "main/internal/requests"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessagesHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, botState *model.BotState) {
	log := config.SetupLogger(config.Config("ENV"))

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "У разраба кривые руки")
	switch botState.CurrentMenu {
	case "main":
		switch update.Message.Text {
		case "Температура":
			msg.Text = "Это температура!"
		case "Розетки":
			msg.Text = "Это розетки!"
		case "Безопасность":
			msg.Text = "Это безопасность!"
		case "Настройки":
			botState.CurrentMenu = "settings"
			KeyboardHandler(bot, update, model.SettingsKeyboard)
			return
		default:
			msg.Text = "Не знаю такого("
		}

	case "settings":
		switch update.Message.Text {
		case "Перезапуск шлюза":
			msg.Text = request.RebootGateway()
		case "Помощь":
			msg.Text = "Это помощь!"
		case "Админ":
			botState.CurrentMenu = "admin"
			KeyboardHandler(bot, update, model.AdminKeyboard)
			return
		case "Вернуться назад":
			botState.CurrentMenu = "main"
			KeyboardHandler(bot, update, model.MainMenuKeyboard)
			return
		default:
			msg.Text = "Не знаю такого("
		}

	case "admin":
		switch update.Message.Text {
		case "Список всех устройств":
			msg.Text = request.PrintAllDevices()
		case "Изменить данные устройства":
			request.ChangeDeviceData(bot)
		case "Вернуться назад":
			botState.CurrentMenu = "settings"
			KeyboardHandler(bot, update, model.SettingsKeyboard)
			return
		default:
			msg.Text = "Не знаю такого("
		}
	}

	if _, err := bot.Send(msg); err != nil {
		log.Error("Error sending the message", "Err", err)
	}
}
