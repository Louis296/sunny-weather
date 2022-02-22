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

func (u *User) GenResp() UserResp {
	return UserResp{
		Name:  u.Name,
		Email: u.Email,
	}
}

type UserResp struct {
	Name  string
	Email string
}
