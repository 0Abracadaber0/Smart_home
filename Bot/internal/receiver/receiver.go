package receiver

import (
	handler "main/internal/handlers"
	model "main/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Receive(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	botState := model.BotState{CurrentMenu: "main"}

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		if update.Message.IsCommand() {
			handler.CommandsHandler(bot, update, &botState)
		} else {
			handler.MessagesHandler(bot, update, &botState)
		}
	}
}
