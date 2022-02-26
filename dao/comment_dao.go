package dao

import (
	"github.com/louis296/sunny-weather/dao/model"
	"gorm.io/gorm"
)

func CreateComment(comment *model.Comment, tx *gorm.DB) error {
	tx = tx.Create(comment)
	return tx.Error
}

func GetCommentById(id int) *model.Comment {
	sql := DB
	var comment model.Comment
	sql = sql.Model(&model.Comment{}).Where("id = ?", id)
	err := sql.First(&comment).Error
	if err != nil {
		return nil
	}
	return &comment
}

func DeleteCommentById(id int, tx *gorm.DB) error {
	tx = tx.Delete(&model.Comment{}, id)
	return tx.Error
}

func DeleteCommentByMomentId(momentId int, tx *gorm.DB) error {
	tx = tx.Delete(&model.Comment{}, "moment_id = ?", momentId)
	return tx.Error
}

func GetCommentsByMomentId(momentId, limit, offset int) ([]*model.Comment, int) {
	sql := DB
	var comments []*model.Comment
	var total int64
	sql = sql.Model(&model.Comment{}).Where("moment_id = ?", momentId)
	sql.Count(&total)
	sql = sql.Limit(limit).Offset(limit * (offset - 1))
	err := sql.Order("create_time desc").Scan(&comments).Error
	if err != nil {
		return nil, 0
	}
	return comments, int(total)
}
