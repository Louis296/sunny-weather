package model

import (
	"gorm.io/gorm"
	"time"
)

type base struct {
	Id         int       `gorm:"id"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
}

func (b base) BeforeCreate(db *gorm.DB) error {
	b.CreateTime = time.Now()
	return nil
}

func (b base) BeforeUpdate(db *gorm.DB) error {
	b.UpdateTime = time.Now()
	return nil
}
