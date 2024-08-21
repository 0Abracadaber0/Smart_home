package request

import (
	"encoding/json"
	"main/config"
	model "main/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
)

func GetAllDevices() []model.Device {
	log := config.SetupLogger(config.Config("ENV"))

	url := config.Config("URL") + "/api/zigbee/devices"

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod("GET")
	req.SetRequestURI(url)

	if err := a.Parse(); err != nil {
		log.Error("Wrong request:", "err", err)
	}

	_, body, err := a.Bytes()
	if err != nil {
		log.Error("Oops!", "err", err)
	}

	var devices []model.Device
	errs := json.Unmarshal(body, &devices)
	if errs != nil {
		log.Error("Oops!", "errs", errs)
	}

	return devices
}

func PrintAllDevices() string {
	devices := GetAllDevices()

	var msg string
	for _, device := range devices {
		msg += device.String()
	}

	return msg
}

func ChangeDeviceData(bot *tgbotapi.BotAPI) {

}
