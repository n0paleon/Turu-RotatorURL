package cmd

import (
	"TuruSMM/delivery/routes"
	"TuruSMM/handler"
	"TuruSMM/internal/config"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func Api() {
	container := dig.New()

	container.Provide(config.NewViper)
	container.Provide(config.NewLogger)
	container.Provide(config.NewFiber)
	container.Provide(config.NewDatabase)
	container.Provide(handler.NewURLHandler)
	container.Provide(routes.NewDeliveryRoutes)

	err := container.Invoke(func(viper *viper.Viper, app *fiber.App, router *routes.DeliveryRoutes, dbConn *gorm.DB) {
		config.Migrator(dbConn, viper)

		app.Use(recover.New())

		app.Use(func(c *fiber.Ctx) error {
			c.Set("X-Author", viper.GetString("app.author"))
			c.Set("X-Served-By", viper.GetString("app.name"))
			c.Set("X-App-Version", viper.GetString("app.version"))

			mode := viper.GetString("app.mode")
			if mode == "maintenance" {
				return c.Status(fiber.StatusServiceUnavailable).SendString("The server is temporarily DOWN due to maintenance!")
			}

			return c.Next()
		}).Name("Copyright Middleware")

		// middleware API monitor page
		app.Use("/monitor", monitor.New(monitor.Config{
			Title:   "Rotator SMM API Status",
			Refresh: 1 * time.Second,
		}))

		router.Register()

		web_host := viper.GetString("web.host")
		web_port := viper.GetInt("web.port")
		app.Listen(fmt.Sprintf("%s:%d", web_host, web_port))
	})

	if err != nil {
		fmt.Printf("error invoke dependency: %+v", err)
	}
}
