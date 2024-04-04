package handler

import (
	"blogv2/config"
	"blogv2/model"
	"blogv2/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type NewCategory struct {
	Name  string  `json:"name"`
	Index *uint64 `json:"index"`
}

type CategoryUpdate struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Index *uint64   `json:"index"`
}

type InputId struct {
	Id uuid.UUID `json:"id"`
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
		return c.Status(fiber.StatusOK).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Failed"})
	}
	return c.Status(fiber.StatusCreated).JSON(&config.Response{Code: 201, Status: true, Message: "Success", Row: row, Data: newCategory})
}

func GetCategory(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	query := new(config.Query)
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	data, pagination, err := repositories.GetListCategory(*query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "failed to get data"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: data, Metadata: pagination})
}

func GetDetailCategory(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	var input InputId
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	category, err := repositories.GetDetailCategoryById(input.Id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Failed to get data"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: category})
}

func UpdateCategory(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	var input CategoryUpdate
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	category, err := repositories.GetCategoryById(input.Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Failed to get data"})
	}
	category.Name = input.Name
	category.Slug = slug.Make(category.Name)
	if input.Index != nil {
		category.Index = *input.Index
	}
	row, err := repositories.UpdateCategory(category)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check data"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Row: row})
}

func DeleteCategory(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	var input InputId
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	row, err := repositories.DeleteCategoryById(input.Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Delete failed"})
	}
	if row < 1 {
		return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: false, Message: "Failed", Row: row})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Row: row})
}
