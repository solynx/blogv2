package database

import (
	"blogv2/model"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Contribute{})
}
