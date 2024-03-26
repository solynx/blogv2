package model

type User struct {
	Model    `gorm:"embedded"`
	Email    string `gorm:"index:,class:FULLTEXT" json:"email" `
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Verified bool   `gorm:"default:false" json:"is_verify"`
}

const (
	ADMIN_ROLE   = "admin"
	DEFAULT_ROLE = "customer"
)
