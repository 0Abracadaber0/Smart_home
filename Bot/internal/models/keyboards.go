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

var SettingsKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Перезапуск шлюза"),
		tgbotapi.NewKeyboardButton("Помощь"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Админ"),
		tgbotapi.NewKeyboardButton("Вернуться назад"),
	),
)

var AdminKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Список всех устройств"),
		tgbotapi.NewKeyboardButton("Изменить данные устройства"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться назад"),
	),
)
