package router

import (
	"blogv2/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	api := app.Group("/api")
	v1 :=  api.Group("/v1")
	v1.Get("/ping", handler.GetPong)
}