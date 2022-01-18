package service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/dao/model"
)

type UserRegisterReq struct {
	Email    string
	Name     string
	Password string
}

type UserRegisterResp struct {
	Name  string
	Email string
}

func (r *UserRegisterReq) Handler(c *gin.Context) (interface{}, error) {
	tx := dao.DB.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	user := model.User{
		Email:    r.Email,
		Name:     r.Name,
		Password: r.Password,
	}
	err := dao.CreateUser(user, tx)
	if err != nil {
		return nil, err
	}
	return UserRegisterResp{
		Name:  r.Name,
		Email: r.Email,
	}, nil
}
