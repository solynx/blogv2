package repositories

import (
	"blogv2/app"
	"blogv2/model"

	"github.com/google/uuid"
)

func CreateUser(user *model.User) error {
	result := app.Core.Database.DB.Create(&user)
	return result.Error
}

func GetUserByUidAndLastLoginId(id, lastLoginId uuid.UUID) (model.User, int64) {
	var user = model.User{}
	result := app.Core.Database.DB.Select("id", "full_name","verified", "metadata", "last_login_id").Where("`id` = ? AND `last_login_id` = ?", id, lastLoginId).Find(&user)
	return user, result.RowsAffected
}

func FindUserLogin(email, password string) (model.User, int64) {
	var user model.User
	result := app.Core.Database.DB.Select("id", "full_name","verified", "metadata").Where("`email` = ? AND `password` = ?", email, password).Find(&user)
	return user, result.RowsAffected
}

func UpdateUser(user model.User) (int64, error) {
	result := app.Core.Database.DB.Updates(&user)
	return result.RowsAffected, result.Error
}
// func GetUserByUid(id uuid.UUID) (model.UserData, int64) {
// 	var userData = model.UserData{}
// 	result := app.Core.Database.DB.Model(&model.User{}).Where("id = ?", id).Find(&userData)
// 	return userData, result.RowsAffected
// }