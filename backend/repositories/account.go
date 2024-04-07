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

func GetUserByUidAndLastLoginId(id uuid.UUID, lastLoginId uuid.UUID, role string) (model.User, int64) {
	var user = model.User{}
	result := app.Core.Database.DB.Select("id", "full_name", "role", "verified", "metadata", "last_login_id").Where("`id` = ? AND `last_login_id` = ? AND `role` = ?", id, lastLoginId, role).Find(&user)
	return user, result.RowsAffected
}

func FindUserLogin(email, password string) (model.User, int64) {
	var user model.User
	result := app.Core.Database.DB.Select("id", "full_name", "role", "verified", "metadata").Where("`email` = ? AND `password` = ?", email, password).Find(&user)
	return user, result.RowsAffected
}

func UpdateUser(user model.User) (int64, error) {
	result := app.Core.Database.DB.Updates(&user)
	return result.RowsAffected, result.Error
}

func GetUserById(id uuid.UUID) (model.User, int64, error) {
	var user model.User
	result := app.Core.Database.DB.Where("id = ?", id).Find(&user) 
	return user, result.RowsAffected, result.Error
} 

