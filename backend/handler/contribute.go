package handler

import (
	"blogv2/config"
	"blogv2/model"
	"blogv2/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ContributeInput struct {
	Content string    `json:"content"`
	PostId  uuid.UUID `json:"post_id"`
	Type    string    `json:"type"`
}

func CreateContributeMessage(c *fiber.Ctx) error {
	var input ContributeInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	if input.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	contribute := model.Contribute{
		Type:    "post",
		Content: input.Content,
		PostId:  input.PostId,
	}
	row, err := repositories.CreateContribute(contribute)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Insert Failed"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Row: row})
}
