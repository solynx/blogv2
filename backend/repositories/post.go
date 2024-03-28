package repositories

import (
	"blogv2/app"
	"blogv2/model"
)

func CreatePost(post *model.Post) (int64, error) {
	result := app.Core.Database.DB.Create(&post)
	return result.RowsAffected, result.Error
}
