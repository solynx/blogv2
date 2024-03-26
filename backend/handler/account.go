package handler

import (
	"blogv2/config"
	"blogv2/model"
	"blogv2/services"

	"crypto/sha256"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

type RegisterSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	return c.SendString("This is login handler")
}

func Register(c *fiber.Ctx) error {
	var userInfo RegisterSchema
	if err := c.BodyParser(&userInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Bad Request"})
	}
	hashedPassword := sha256.Sum256([]byte(userInfo.Password))
	var newUser = model.User{
		Email:    userInfo.Email,
		Password: hex.EncodeToString(hashedPassword[:]),
		Role:     model.ADMIN_ROLE,
	}
	err := services.CreateUser(&newUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Register failed!"})
	}
	return c.Status(fiber.StatusCreated).JSON(&config.ErrorSchema{Code: 200, Status: true, Message: "Create user successfully!"})
}
