package model

type ThumbUp struct {
	Base
	MomentId int `gorm:"moment_id"`
	UserId   int `gorm:"user_id"`
}

func (t *ThumbUp) TableName() string {
	return "thumb_up"
}
