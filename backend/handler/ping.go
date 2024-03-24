package handler

import "github.com/gofiber/fiber/v2"

func GetPong(c *fiber.Ctx) error {
	return c.SendString("PONG")
}