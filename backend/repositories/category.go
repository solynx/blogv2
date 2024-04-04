package repositories

import (
	"blogv2/app"
	"blogv2/config"
	"blogv2/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateCategory(category *model.Category) (int64, error) {
	result := app.Core.Database.DB.Create(&category)
	return result.RowsAffected, result.Error
}

func GetListCategory(query config.Query) ([]*model.Category, config.Pagination, error) {
	var categories []*model.Category
	var pagination config.Pagination
	tx := app.Core.Database.DB.Model(&model.Category{})
	tx = tx.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("`id`, `full_name`")
	})

	tx.Offset(query.GetOffset()).Limit(query.Limit).Order("`created_at` " + query.GetSort()).Select("id, name, user_id, created_at").Find(&categories)
	tx.Count(&pagination.Total)
	pagination.Page = query.GetPage()
	return categories, pagination, tx.Error
}

func GetCategoryById(id uuid.UUID) (model.Category, error) {
	var category model.Category
	result := app.Core.Database.DB.First(&category, "id = ?", id)
	return category, result.Error
}

func UpdateCategory(category model.Category) (int64, error) {
	result := app.Core.Database.DB.Debug().Updates(category)
	return result.RowsAffected, result.Error
}

func DeleteCategoryById(id uuid.UUID) (int64, error) {
	result := app.Core.Database.DB.Where("id = ?", id).Delete(&model.Category{})
	return result.RowsAffected, result.Error
}

func GetDetailCategoryById(id uuid.UUID) (*model.Category, error) {
	var category model.Category
	tx := app.Core.Database.DB.Model(&model.Category{})
	tx = tx.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("`id`, `full_name`")
	})
	result := tx.First(&category, "id = ?", id)
	return &category, result.Error
}
