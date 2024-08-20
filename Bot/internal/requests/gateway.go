package request

import (
	"encoding/json"
	"main/config"

	"github.com/gofiber/fiber/v2"
)

func RebootGateway() string {
	log := config.SetupLogger(config.Config("ENV"))

	url := config.Config("URL") + "/api/reboot"

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

	if response["success"].(bool) == true {
		return "Шлюз успешно перезапущен!"
	} else {
		return "Ошибка при перезапуске шлюза!"
	}

}
