package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	Model    `gorm:"embedded"`
	Email    string          `gorm:"uniqueIndex" json:"email"`
	Password string          `json:"password"`
	FullName string          `json:"full_name"`
	Role     string          `json:"role"`
	Verified bool            `gorm:"default:false" json:"is_verify"`
	Metadata *datatypes.JSON `json:"metadata"`
}

type UserData struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
	Verified bool      `json:"is_verify"`
	ImageUrl string    `json:"image_url"`
}

const (
	ADMIN_ROLE   = "admin"
	DEFAULT_ROLE = "customer"
)
