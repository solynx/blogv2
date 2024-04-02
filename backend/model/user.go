package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	Model       `gorm:"embedded"`
	Email       string          `gorm:"uniqueIndex" json:"email,omitempty"`
	Password    string          `json:"password,omitempty"`
	FullName    string          `json:"full_name,omitempty"`
	Role        string          `json:"role,omitempty"`
	Verified    bool            `gorm:"default:false" json:"is_verify,omitempty"`
	LastLoginId uuid.UUID       `gorm:"default:NULL" json:"-"`
	Metadata    *datatypes.JSON `json:"metadata,omitempty"`
}

type UserPayload struct {
	ID          uuid.UUID `json:"id"`
	LastLoginId uuid.UUID `json:"login_id"`
}

const (
	ADMIN_ROLE   = "admin"
	DEFAULT_ROLE = "customer"
)
