package request

import (
	"encoding/json"
	"fmt"
	"main/config"
	model "main/internal/models"
	"net/url"
	"reflect"

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

func ChangeDeviceData(bot *tgbotapi.BotAPI, update tgbotapi.Update) string {
	log := config.SetupLogger(config.Config("ENV"))

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите nwkAddr:")
	if _, err := bot.Send(msg); err != nil {
		log.Error(err.Error())
	}

	userId := update.Message.From.ID
	var nwkAddr string
	var foundDevice model.Device
	var found bool

	for update := range bot.GetUpdatesChan(tgbotapi.NewUpdate(0)) {
		if update.Message.From.ID == userId && update.Message.Text != "" {
			if nwkAddr == "" {
				nwkAddr = update.Message.Text
				foundDevice, found = findDevice("NwkAddr", nwkAddr)
				if found {
					msg := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						fmt.Sprintf(
							"Старое имя: %s\nВведите новое имя:",
							foundDevice.Friendly_name,
						),
					)
					if _, err := bot.Send(msg); err != nil {
						log.Error(err.Error())
					}
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такой девайс не найден!")
					if _, err := bot.Send(msg); err != nil {
						log.Error(err.Error())
					}

					return "Возникла ошибка при попытке изменения имени"
				}
			} else {
				newName := update.Message.Text
				url := config.Config("URL") + fmt.Sprintf(
					"/api/zigbee/rename?old=%s&new=%s",
					url.QueryEscape(foundDevice.Friendly_name),
					url.QueryEscape(newName),
				)

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

				var response fiber.Map
				errs := json.Unmarshal(body, &response)
				if errs != nil {
					log.Error("Oops!", "errs", errs)
				}

				if response["success"].(bool) {
					return "Имя успешно изменено"
				} else {
					return "Ошибка при изменении имени!"
				}
			}
		}
	}

	return "Возникла ошибка при попытке изменения имени"
}

func findDevice(searchParam string, searchValue interface{}) (model.Device, bool) {
	devices := GetAllDevices()
	for _, device := range devices {
		value := reflect.ValueOf(device).FieldByName(searchParam).Interface()

		if reflect.DeepEqual(value, searchValue) {
			return device, true
		}
	}

	return model.Device{}, false
}
