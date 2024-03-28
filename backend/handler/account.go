package handler

import (
	"blogv2/config"
	"blogv2/model"
	repo "blogv2/repositories"

	"crypto/sha256"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

type Identity struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var userInfo Identity
	if err := c.BodyParser(&userInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request"})
	}
	hashedPassword := sha256.Sum256([]byte(userInfo.Password))
	userQuery := model.User{
		Email:    userInfo.Email,
		Password: hex.EncodeToString(hashedPassword[:]),
	}
	user, row := repo.GetUser(&userQuery)
	if row == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: user})
}

func Register(c *fiber.Ctx) error {
	var userInfo Identity
	if err := c.BodyParser(&userInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request"})
	}
	hashedPassword := sha256.Sum256([]byte(userInfo.Password))
	var newUser = model.User{
		Email:    userInfo.Email,
		Password: hex.EncodeToString(hashedPassword[:]),
		Role:     model.ADMIN_ROLE,
	}
	err := repo.CreateUser(&newUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Register failed!"})
	}
	return c.Status(fiber.StatusCreated).JSON(&config.Response{Code: 200, Status: true, Message: "Create user successfully!"})
}
