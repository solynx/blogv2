package model

import "github.com/google/uuid"

type Post struct {
	Model          `gorm:"embedded"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Slug           string    `json:"slug"`
	Content        string    `gorm:"type:text" json:"content"`
	UserId         uuid.UUID `gorm:"type:text" json:"user_id"`
	User           User      `gorm:"foreignKey:UserId"`
	SEOTitle       string    `json:"seo_title"`
	SEOKeyword     string    `json:"seo_keyword"`
	SEODescription string    `json:"seo_description"`
}
