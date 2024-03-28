package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Category struct {
	Model    `gorm:"embedded"`
	Name     string          `json:"name"`
	UserId   uuid.UUID       `json:"user_id"`
	Index    uint64          `gorm:"default:1" json:"index"`
	Metadata *datatypes.JSON `json:"metadata"`
	User     User            `gorm:"foreignKey:UserId"`
}
