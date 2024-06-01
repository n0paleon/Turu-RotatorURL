package handler

import (
	"math/rand"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/gofiber/fiber/v2"
	ua2 "github.com/mileusna/useragent"
	"github.com/mssola/useragent"
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

	ect := []string{"4g", "5g", "3g", "4g", "4g", "4g", "4g", "4g", "5g", "5g"}
	downlink := []float64{1.5, 3, 5, 10, 20, 30, 40, 50}
	ram := []int{1, 2, 3, 4, 6, 8, 10, 12, 16, 20, 24, 32, 64}

	detail := useragent.New(ua)
	detail.Parse(ua)

	rand.Seed(time.Now().UnixNano())

	info := ua2.Parse(ua)

	platform := detail.OSInfo().Name
	version := detail.OSInfo().Version
	if len(platform) < 1 {
		platform = "Windows"
		version = "10.0.0"
	}
	is_mobile := detail.Mobile()
	vp_width := 1280
	vp_height := 720

	is_mobile_str := "?0"
	if is_mobile {
		vp_width = 720
		vp_height = 1440
		is_mobile_str = "?1"
	}

	return ctx.JSON(&fiber.Map{
		"result": &fiber.Map{
			"ua":          ua,
			"ua_version":  info.Version,
			"is_mobile":   is_mobile_str,
			"os_platform": platform,
			"os_model":    detail.Model(),
			"os_version":  version,
			"viewport": &fiber.Map{
				"width":  vp_width,
				"height": vp_height,
			},
			"connectivity":  ect[rand.Intn(len(ect))],
			"downlink":      downlink[rand.Intn(len(downlink))],
			"device_memory": ram[rand.Intn(len(ram))],
			"prefer_color":  []string{"dark", "light", "no-preference"}[rand.Intn(3)],
			"arch":          []string{"x86", "ARM"}[rand.Intn(2)],
		},
	})
}
