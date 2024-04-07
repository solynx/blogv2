package handler

import (
	"blogv2/app"
	"blogv2/config"
	"blogv2/model"
	"blogv2/repositories"
	"log"
	"time"

	"crypto/sha256"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Identity struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PayloadClaims struct {
	Payload model.UserPayload `json:"payload"`
	jwt.RegisteredClaims
}

type UpdateInput struct{
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var userInfo Identity
	if err := c.BodyParser(&userInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request"})
	}
	hashedPassword := sha256.Sum256([]byte(userInfo.Password))

	user, row := repositories.FindUserLogin(userInfo.Email, hex.EncodeToString(hashedPassword[:]))
	if row == 0 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	//update last login
	user.LastLoginId = uuid.New()
	row, err := repositories.UpdateUser(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Please check connect"})
	}
	if row < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Please check id"})
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"user_id":       user.ID,
		"last_login_id": user.LastLoginId,
		"role":          user.Role,
		"exp":           time.Now().Add(time.Hour * 168).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString(app.Core.Server.RouterPrivateKey)
	if err != nil {
		log.Fatal(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: user, Token: t})
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
	err := repositories.CreateUser(&newUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Register failed!"})
	}
	return c.Status(fiber.StatusCreated).JSON(&config.Response{Code: 200, Status: true, Message: "Create user successfully!"})
}

func GetProfileUser(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Permission denied"})
	} 
	user, row, err := repositories.GetUserById(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Can't find data"})
	}
	if row == 0 {
		return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 400, Status: false, Message: "Can't find data"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: user})
}

func UpdateProfileUser(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Permission denied"})
	} 
	var input UpdateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Check your arg"})
	}
	user, row, err := repositories.GetUserById(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Can't find data"})
	}
	if row == 0 {
		return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 400, Status: false, Message: "Can't find data"})
	}
	if len(input.Password) > 8 {
		hashedPassword := sha256.Sum256([]byte(input.Password))
		user.Password = hex.EncodeToString(hashedPassword[:])
	}
	if input.FullName != "" {
		user.FullName = input.FullName
	}
	_, errU := repositories.UpdateUser(user)
	if errU != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Update failed"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success"})
}
