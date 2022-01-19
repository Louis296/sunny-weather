package dao

import (
	"github.com/louis296/sunny-weather/dao/model"
	"gorm.io/gorm"
)

func CreateUser(user model.User, tx *gorm.DB) error {
	tx = tx.Create(&user)
	return tx.Error
}

func GetUserByEmail(email string) *model.User {
	sql := DB
	var user model.User
	sql = sql.Model(&model.User{}).Where("email=?", email)
	err := sql.First(&user).Error
	if err != nil {
		return nil
	}
	return &user
}
