package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Category struct {
	Model    `gorm:"embedded"`
	Name     string          `json:"name"`
	Slug     string          `gorm:"uniqueIndex" json:"slug,omitempty"`
	UserId   uuid.UUID       `json:"user_id"`
	Index    uint64          `gorm:"default:99" json:"index,omitempty"`
	Metadata *datatypes.JSON `json:"metadata,omitempty"`
	User     *User           `gorm:"foreignKey:UserId" json:"user,omitempty"`
}
