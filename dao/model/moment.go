package model

type Moment struct {
	Base
	Context string `gorm:"context"`
	Like    int    `gorm:"like"`
	UserId  int    `gorm:"user_id"`
}

func (m *Moment) TableName() string {
	return "moment"
}
