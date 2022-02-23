package model

import "time"

type Comment struct {
	Base
	Context  string `gorm:"context"`
	UserId   int    `gorm:"user_id"`
	MomentId int    `gorm:"moment_id"`
}

func (c *Comment) TableName() string {
	return "comment"
}

type CommentResp struct {
	Id         int
	Context    string
	User       UserResp
	CreateTime time.Time
	UpdateTime time.Time
}

func (c *Comment) GenResp() CommentResp {
	return CommentResp{
		Id:         c.Id,
		Context:    c.Context,
		CreateTime: c.CreateTime,
		UpdateTime: c.UpdateTime,
	}
}
