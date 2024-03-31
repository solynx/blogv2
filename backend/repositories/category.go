package repositories

import (
	"blogv2/app"
	"blogv2/model"
)

func CreateCategory(category *model.Category) (int64, error) {
	result := app.Core.Database.DB.Create(&category)
	return result.RowsAffected, result.Error
}