package repositories

import (
	"blogv2/app"
	"blogv2/model"
)

func CreateContribute(contribute model.Contribute) (int64, error) {
	result := app.Core.Database.DB.Create(&contribute)
	return result.RowsAffected, result.Error
}
