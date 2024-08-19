package main

import (
	"main/config"
	"main/internal/connect"
	"main/internal/receiver"
)

func main() {
	log := config.SetupLogger(config.Config("ENV"))
	log.Info("starting server", "env", config.Config("ENV"))

	bot := connect.ConnectBot()

	receiver.Receive(bot)

}
