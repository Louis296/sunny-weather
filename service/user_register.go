package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/dao/model"
	"github.com/louis296/sunny-weather/pkg/common"
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
	existUser := dao.GetUserByEmail(r.Email)
	if existUser != nil {
		return nil, errors.New("duplicate user email")
	}
	user := model.User{
		Email:    r.Email,
		Name:     r.Name,
		Password: common.GetMD5WithSalt(r.Password),
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
