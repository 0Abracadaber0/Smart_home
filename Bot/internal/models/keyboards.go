package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MainMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Температура"),
		tgbotapi.NewKeyboardButton("Розетки"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Безопасность"),
		tgbotapi.NewKeyboardButton("Настройки"),
	),
)
