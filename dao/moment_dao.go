package dao

import (
	"github.com/louis296/sunny-weather/dao/model"
	"gorm.io/gorm"
)

func CreateMoment(moment *model.Moment, tx *gorm.DB) error {
	tx = tx.Create(moment)
	return tx.Error
}

func GetMomentById(id int) *model.Moment {
	sql := DB
	var moment model.Moment
	sql = sql.Model(&model.Moment{}).Where("id=?", id)
	err := sql.First(&moment).Error
	if err != nil {
		return nil
	}
	return &moment
}

func GetMomentsByUserId(userId int) []*model.Moment {
	sql := DB
	var moments []*model.Moment
	sql = sql.Model(&model.Moment{}).Where("user_id=?", userId)
	err := sql.Scan(&moments).Error
	if err != nil {
		return nil
	}
	return moments
}
