package model

type Post struct {
	Model       `gorm:"embedded"`
	Title       string
	Description string
	Content     string
	SEOTitle    string
}
