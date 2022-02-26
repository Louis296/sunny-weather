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

func GetMomentsByUserId(userId, limit, offset int) ([]*model.Moment, int) {
	sql := DB
	var moments []*model.Moment
	var total int64
	sql = sql.Model(&model.Moment{}).Where("user_id=?", userId)
	sql.Count(&total)
	sql = sql.Limit(limit).Offset(limit * (offset - 1))
	err := sql.Order("update_time desc").Scan(&moments).Error
	if err != nil {
		return nil, 0
	}
	return moments, int(total)
}

func GetAllMoments(limit, offset int) ([]*model.Moment, int) {
	sql := DB
	var moments []*model.Moment
	var total int64
	sql = sql.Model(&model.Moment{})
	sql.Count(&total)
	sql = sql.Limit(limit).Offset(limit * (offset - 1))
	err := sql.Order("update_time desc").Scan(&moments).Error
	if err != nil {
		return nil, 0
	}
	return moments, int(total)
}

func DeleteMomentById(id int, tx *gorm.DB) error {
	tx = tx.Delete(&model.Moment{}, id)
	return tx.Error
}

func SaveMoment(moment *model.Moment, tx *gorm.DB) error {
	tx = tx.Save(moment)
	return tx.Error
}

func GetMomentUserThumb(momentId, userId int) *model.ThumbUp {
	sql := DB
	sql = sql.Model(&model.ThumbUp{}).Where("moment_id = ? and user_id = ?", momentId, userId)
	var thumbUp model.ThumbUp
	if err := sql.First(&thumbUp).Error; err != nil {
		return nil
	}
	return &thumbUp
}

func CreateMomentThumbUp(thumbUp *model.ThumbUp, tx *gorm.DB) error {
	tx = tx.Create(thumbUp)
	return tx.Error
}
