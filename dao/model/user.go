package model

type User struct {
	Base
	Email    string `gorm:"email"`
	Name     string `gorm:"name"`
	Password string `gorm:"password"`
}

func (u *User) TableName() string {
	return "user"
}
