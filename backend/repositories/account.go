package repositories

import (
	"blogv2/app"
	"blogv2/model"
)

func CreateUser(user *model.User) error {
	result := app.Core.Database.DB.Create(&user)
	return result.Error
}

func GetUser(user *model.User) (model.UserData, int64) {
	var userData = model.UserData{}
	result := app.Core.Database.DB.Model(&model.User{}).Where(&user).Find(&userData)
	return userData, result.RowsAffected
}
