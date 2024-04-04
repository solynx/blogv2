package repositories

import (
	"blogv2/app"
	"blogv2/config"
	"blogv2/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreatePost(post *model.Post) (int64, error) {
	result := app.Core.Database.DB.Create(&post)
	return result.RowsAffected, result.Error
}

func GetListPost(query config.Query) ([]*model.Post, config.Pagination, error) {
	var posts []*model.Post
	var pagination config.Pagination
	tx := app.Core.Database.DB.Model(&model.Post{})
	tx = tx.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("`id`, `full_name`")
	})
	tx = tx.Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("`id`, `name`")
	})
	tx.Offset(query.GetOffset()).Limit(query.Limit).Order("`created_at` " + query.GetSort()).Select("id, title, user_id, category_id, created_at").Find(&posts)
	tx.Count(&pagination.Total)
	pagination.Page = query.GetPage()
	return posts, pagination, tx.Error
}

func GetPostById(id uuid.UUID) (model.Post, error) {
	var post model.Post
	tx := app.Core.Database.DB.Where("id = ?", id).Omit("updated_at")
	tx = tx.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("`id`, `full_name`")
	})
	tx = tx.Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("`id`, `name`")
	})
	tx.Find(&post)
	return post, tx.Error
}

func UpdatePost(post *model.Post) (int64, error) {
	result := app.Core.Database.DB.Debug().Updates(&post)
	return result.RowsAffected, result.Error
}

func DeletePost(id uuid.UUID) (int64, error) {
	result := app.Core.Database.DB.Where("id = ?", id).Delete(&model.Post{})
	return result.RowsAffected, result.Error
}
