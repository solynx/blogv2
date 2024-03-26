package services

import (
	"blogv2/app"
	"blogv2/model"
)

func CreateUser(user *model.User) error {
	result := app.Core.Database.DB.Create(&user)
	return result.Error
}
