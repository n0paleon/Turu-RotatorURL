package routes

import (
	"TuruSMM/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DeliveryRoutes struct {
	App           *fiber.App
	Log           *logrus.Logger
	URLHandler    *handler.URLHandler
	ClientHandler *handler.ClientHandler
}

func (route *DeliveryRoutes) Register() {
	r := route.App.Group("/api/v1")

	r.Get("/rotator", route.URLHandler.Rotator)
	r.Get("/rotator/:id", route.URLHandler.RotateByID)
	r.Get("/rotate/:id", route.URLHandler.CustomRotate)

	r.Get("/client/ua", route.ClientHandler.RandomUA)
}

func NewDeliveryRoutes(App *fiber.App, Log *logrus.Logger, URLHandler *handler.URLHandler) *DeliveryRoutes {
	return &DeliveryRoutes{
		App:        App,
		Log:        Log,
		URLHandler: URLHandler,
	}
}
