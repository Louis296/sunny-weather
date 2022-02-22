package model

type Comment struct {
	Base
	Context  string `gorm:"context"`
	UserId   int    `gorm:"user_id"`
	MomentId int    `gorm:"moment_id"`
}

func (c *Comment) TableName() string {
	return "comment"
}
