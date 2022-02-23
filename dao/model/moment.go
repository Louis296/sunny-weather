package model

import "time"

type Moment struct {
	Base
	Context string `gorm:"context"`
	Like    int    `gorm:"like"`
	UserId  int    `gorm:"user_id"`
}

func (m *Moment) TableName() string {
	return "moment"
}

type MomentResp struct {
	Id         int
	Context    string
	Like       int
	UpdateTime time.Time
	CreateTime time.Time
	User       UserResp
}

func (m *Moment) GenResp() MomentResp {
	return MomentResp{
		Id:         m.Id,
		Context:    m.Context,
		Like:       m.Like,
		UpdateTime: m.UpdateTime,
		CreateTime: m.CreateTime,
	}
}
