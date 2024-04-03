package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Post struct {
	Model          `gorm:"embedded"`
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	Slug           string          `gorm:"uniqueIndex" json:"slug"`
	Content        string          `gorm:"type:text" json:"content"`
	UserId         uuid.UUID       `gorm:"type:text" json:"user_id"`
	CategoryId     uuid.UUID       `json:"category_id"`
	User           User            `gorm:"foreignKey:UserId" json:"user"`
	Category       Category        `gorm:"foreignKey:CategoryId" json:"category"`
	Metadata       *datatypes.JSON `json:"metadata"`
	SEOTitle       *string         `json:"seo_title"`
	SEOKeyword     *string         `json:"seo_keyword"`
	SEODescription *string         `json:"seo_description"`
}
