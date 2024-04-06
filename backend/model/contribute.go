package model

import "github.com/google/uuid"

type Contribute struct {
	Model   `gorm:"embedded"`
	Content string    `gorm:"type:text" json:"content,omitempty"`
	PostId  uuid.UUID `json:"-"`
	Post    *Post     `gorm:"foreignKey:PostId" json:"post"`
	Type    string    `json:"-"`
}
