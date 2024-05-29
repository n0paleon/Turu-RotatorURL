package config

import (
	"time"

	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper, log *logrus.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:           config.GetString("app.name"),
		CaseSensitive:     true,
		EnablePrintRoutes: true,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		Prefork:           config.GetBool("web.prefork"),
		StrictRouting:     true,
		WriteTimeout:      30 * time.Second,
		ErrorHandler:      NewErrorHandler(config, log),
	})

	return app
}

func NewErrorHandler(config *viper.Viper, log *logrus.Logger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		log.Warnf("fiber returned an error: %+v", err)

		app_mode := config.GetString("app.mode")
		if app_mode == "production" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "error while processing your request",
			})
		}

		return ctx.Status(code).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
}
