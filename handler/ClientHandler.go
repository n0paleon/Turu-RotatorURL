package handler

import (
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ClientHandler struct {
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewClientHandler(log *logrus.Logger, config *viper.Viper) *ClientHandler {
	return &ClientHandler{
		Log:    log,
		Config: config,
	}
}

func (h *ClientHandler) RandomUA(ctx *fiber.Ctx) error {
	client := browser.Client{
		MaxPage: 3,
		Delay:   200 * time.Millisecond,
		Timeout: 10 * time.Second,
	}
	cache := browser.Cache{}
	b := browser.NewBrowser(client, cache)

	var ua string

	switch ctx.Query("type") {
	case "mobile":
		ua = b.Mobile()
	case "desktop":
		ua = b.Computer()
	case "ios":
		ua = b.IOS()
	case "android":
		ua = b.Android()
	default:
		ua = b.Random()
	}

	return ctx.JSON(&fiber.Map{
		"result": ua,
	})
}
