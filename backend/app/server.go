package app

import (
	"github.com/gofiber/fiber/v2"
)

type ServerConfig struct {
	App  *fiber.App
	Port string
}

func (server *ServerConfig) Start() {
	server.App.Listen(server.Port)
}

func (server *ServerConfig) Setup() {
	server.App = fiber.New()
	server.Port = ":8080"
}
