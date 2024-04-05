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
	Content        string          `gorm:"type:text" json:"content,omitempty"`
	UserId         uuid.UUID       `json:"-"`
	CategoryId     uuid.UUID       `json:"-"`
	Published      *bool           `gorm:"type:boolean;column:published;default:false" json:"published,omitempty"`
	User           *User           `gorm:"foreignKey:UserId" json:"user"`
	Category       *Category       `gorm:"foreignKey:CategoryId" json:"category"`
	Metadata       *datatypes.JSON `json:"metadata,omitempty"`
	SEOTitle       *string         `json:"seo_title,omitempty"`
	SEOKeywords    *string         `json:"seo_keywords,omitempty"`
	SEODescription *string         `json:"seo_description,omitempty"`
}
