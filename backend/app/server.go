package app

import (
	"blogv2/router"

	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	App *fiber.App
	Port string
}


func (server *ServerConfig) Start() {
	server.Setup()
	server.SetupRouter()
	server.App.Listen(server.Port)
}

func (server *ServerConfig) Setup() {
	server.App = fiber.New()
	server.Port = ":8080"
}

func (server *ServerConfig) SetupRouter() {
	router.SetupRouter(server.App)
}