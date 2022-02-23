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

func GetUserById(id int) *model.User {
	sql := DB
	var user model.User
	sql = sql.Model(&model.User{}).Where("id=?", id)
	err := sql.First(&user).Error
	if err != nil {
		return nil
	}
	return &user
}

func GetUserIdLookup(ids ...int) (map[int]model.User, error) {
	var users []model.User
	sql := DB
	sql = sql.Model(&model.User{}).Where("id in ?", ids)
	if err := sql.Scan(&users).Error; err != nil {
		return nil, err
	}
	result := make(map[int]model.User)
	for _, user := range users {
		result[user.Id] = user
	}
	return result, nil
}
