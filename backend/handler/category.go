package handler

import (
	"blogv2/config"
	"blogv2/model"
	"blogv2/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type NewCategory struct {
	Name string `json:"name"`
	Index *uint64 `json:"index"`
}

func CreateCategory(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	var body NewCategory
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check your arg"})
	}
	if body.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check your arg"})
	}
	newCategory := &model.Category{Name: body.Name, UserId: user.ID, Slug: slug.Make(body.Name)}
	if body.Index != nil {
		newCategory.Index = *body.Index
	}
	row, err := repositories.CreateCategory(newCategory)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Failed"})
	}
	return c.Status(fiber.StatusCreated).JSON(&config.Response{Code: 201, Status: true, Message: "Success", Row: row, Data: newCategory})
}