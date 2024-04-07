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
	Content        string    `json:"content"`
	Published *bool 	`json:"published"`
	SEOTitle       *string   `json:"seo_title"`
	SEOKeywords    *string   `json:"seo_keywords"`
	SEODescription *string   `json:"seo_description"`
	CategoryId     uuid.UUID `json:"category_id"`
}

type PostUpdateInput struct {
	Id             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Content        string    `gorm:"type:text" json:"content"`
	SEOTitle       *string   `json:"seo_title"`
	SEOKeywords    *string   `json:"seo_keywords"`
	SEODescription *string   `json:"seo_description"`
	CategoryId     uuid.UUID `json:"category_id"`
	Published      bool      `json:"published"`
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
	newPost := &model.Post{Title: args.Title, UserId: user.ID, Published: args.Published, CategoryId: args.CategoryId, Description: args.Description, Content: args.Content, SEOTitle: args.SEOTitle, SEOKeywords: args.SEOKeywords, SEODescription: args.SEODescription}
	newPost.ID = uuid.New()
	newPost.Slug = helpers.GetSlug(newPost.Title, newPost.ID)
	row, err := repositories.CreatePost(newPost)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 404, Status: false, Message: "Create failed!"})
	}
	return c.JSON(&config.Response{Code: 200, Message: "Create success!", Status: true, Row: row})
}

func GetListPost(c *fiber.Ctx) error {
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
	post.SEOKeywords = input.SEOKeywords
	post.SEOTitle = input.SEOTitle
	post.Published = &input.Published
	post.Category = nil
	post.User = nil
	row, err := repositories.UpdatePost(&post)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 400, Status: false, Message: "Bad Request", Error: "Failed to update"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Row: row, Data: post})
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

func GetPostDetail(c *fiber.Ctx) error {
	_, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please login"})
	}
	query := new(config.Query)
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	post, err := repositories.GetPostById(query.Id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check your arg"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: post})
}

//PUBLIC API FOR UI

func GetPostDetailBySlug(c *fiber.Ctx) error {
	query := new(config.UIQuery)
	if err := c.QueryParser(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	if query.Slug == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	post, err := repositories.GetPostBySlug(query.Slug)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Can't find data"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: post})
}

func PublicGetListPost(c *fiber.Ctx) error {
	var uiQuery config.UIQuery
	if err := c.BodyParser(&uiQuery); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	data, pagination, err := repositories.GetListPostForUI(uiQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Can't get data"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: data, Metadata: pagination})
}

func PublicGetRelatedPosts(c *fiber.Ctx) error {
	//lay cac bai viet lien quan den author hoac category
	var uiQuery config.UIQuery
	if err := c.BodyParser(&uiQuery); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	data, err := repositories.GetListRelatedPost(uiQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Can't get data"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: data})
}

func PublicGetListPostByCategory(c *fiber.Ctx) error {
	var uiQuery config.UIQuery
	if err := c.BodyParser(&uiQuery); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Please check arg"})
	}
	data, pagination, err := repositories.GetListPostForUIByCategoryId(uiQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorSchema{Code: 400, Status: false, Message: "Can't get data"})
	}
	return c.Status(fiber.StatusOK).JSON(&config.Response{Code: 200, Status: true, Message: "Success", Data: data, Metadata: pagination})
}
