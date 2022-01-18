package dao

import (
	"github.com/louis296/sunny-weather/dao/model"
	"gorm.io/gorm"
)

func CreateUser(user model.User, tx *gorm.DB) error {
	tx = tx.Create(&user)
	return tx.Error
}
