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
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Content        string  `gorm:"type:text" json:"content"`
	SEOTitle       *string `json:"seo_title"`
	SEOKeyword     *string `json:"seo_keyword"`
	SEODescription *string `json:"seo_description"`
}

func CreatePost(c *fiber.Ctx) error {
	var args PostSchema
	if err := c.BodyParser(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request", Error: "Please check args sended"})
	}
	newPost := &model.Post{Title: args.Title, UserId: uuid.MustParse("9883e1f5-bab9-42bf-af2e-34e0a3df0a6d"), Description: args.Description, Content: args.Content, SEOTitle: args.SEOTitle, SEOKeyword: args.SEOKeyword, SEODescription: args.SEODescription}
	newPost.ID = uuid.New()
	newPost.Slug = helpers.GetSlug(newPost.Title, newPost.ID)
	row, err := repositories.CreatePost(newPost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.Response{Code: 404, Status: false, Message: "Create failed!"})
	}
	return c.JSON(&config.Response{Code: 200, Message: "Create success!", Status: true, Row: row})
}
