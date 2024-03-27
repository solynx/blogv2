package model

import "github.com/google/uuid"

type User struct {
	Model    `gorm:"embedded"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Verified bool   `gorm:"default:false" json:"is_verify"`
}

type UserData struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `gorm:"uniqueIndex" json:"email"`
	FullName string    `json:"full_name"`
	Verified bool      `gorm:"default:false" json:"is_verify"`
}

const (
	ADMIN_ROLE   = "admin"
	DEFAULT_ROLE = "customer"
)
