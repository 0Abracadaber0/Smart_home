package main

import (
	"main/config"
	handler "main/internal/Handlers"
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
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		switch update.Message.Text {
		case "menu":
			handler.MainMenuHandler(bot, update)
		}

	}

}
