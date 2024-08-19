package connect

import (
	"fmt"
	"main/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ConnectBot() *tgbotapi.BotAPI {
	log := config.SetupLogger(config.Config("ENV"))

	bot, err := tgbotapi.NewBotAPI(config.Config("BOT_TOKEN"))
	if err != nil {
		log.Error("An error occurred while connecting to the bot", "Error text", err)
	}

	bot.Debug = true
	log.Info(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	return bot
}
