package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Category struct {
	Model    `gorm:"embedded"`
	Name     string          `json:"name"`
	Slug           string          `gorm:"uniqueIndex" json:"slug"`
	UserId   uuid.UUID       `json:"user_id"`
	Index    uint64          `gorm:"default:99" json:"index"`
	Metadata *datatypes.JSON `json:"metadata"`
	User     *User            `gorm:"foreignKey:UserId" json:"user,omitempty"`
}
