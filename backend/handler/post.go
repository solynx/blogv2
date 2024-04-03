package handler

import (
	"blogv2/config"
	"blogv2/helpers"
	"blogv2/model"
	"blogv2/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PostSchema struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Content        string    `gorm:"type:text" json:"content"`
	SEOTitle       *string   `json:"seo_title"`
	SEOKeyword     *string   `json:"seo_keyword"`
	SEODescription *string   `json:"seo_description"`
	CategoryId     uuid.UUID `json:"category_id"`
}

type PostUpdateInput struct {
	Id             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Content        string    `gorm:"type:text" json:"content"`
	SEOTitle       *string   `json:"seo_title"`
	SEOKeyword     *string   `json:"seo_keyword"`
	SEODescription *string   `json:"seo_description"`
	CategoryId     uuid.UUID `json:"category_id"`
}

type IdInput struct {
	Id uuid.UUID `json:"id"`
}

func CreatePost(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	var args PostSchema
	if err := c.BodyParser(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request", Error: "Please check args sended"})
	}
	newPost := &model.Post{Title: args.Title, UserId: user.ID, CategoryId: args.CategoryId, Description: args.Description, Content: args.Content, SEOTitle: args.SEOTitle, SEOKeyword: args.SEOKeyword, SEODescription: args.SEODescription}
	newPost.ID = uuid.New()
	newPost.Slug = helpers.GetSlug(newPost.Title, newPost.ID)
	row, err := repositories.CreatePost(newPost)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 404, Status: false, Message: "Create failed!"})
	}
	return c.JSON(&config.Response{Code: 200, Message: "Create success!", Status: true, Row: row})
}

func GetPost(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	query := new(config.Query)
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	data, pagination, err := repositories.GetListPost(*query)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "failed to get data"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: data, Metadata: pagination})
}

func UpdatePost(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	var input PostUpdateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request", Error: "Please check args sended"})
	}
	post, err := repositories.GetPostById(input.Id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request", Error: "Can't find record"})
	}
	post.Title = input.Title
	post.Slug = helpers.GetSlug(post.Title, post.ID)
	post.CategoryId = input.CategoryId
	post.Description = input.Description
	post.Content = input.Content
	post.SEODescription = input.SEODescription
	post.SEOKeyword = input.SEOKeyword
	post.SEOTitle = input.SEOTitle
	row, err := repositories.UpdatePost(*post)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request", Error: "Failed to update"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Row: row})
}

func DeletePost(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	var input InputId
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	row, err := repositories.DeletePost(input.Id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request", Error: "Failed to delete"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Row: row})
}
