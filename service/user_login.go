package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/pkg/common"
	"github.com/louis296/sunny-weather/pkg/jwt"
)

type UserLoginReq struct {
	Email    string
	Password string
}

type UserLoginResp struct {
	Token string
}

func (r *UserLoginReq) Handler(c *gin.Context) (interface{}, error) {
	user := dao.GetUserByEmail(r.Email)
	if user == nil {
		return nil, errors.New("no such user")
	}
	if common.GetMD5WithSalt(r.Password) != user.Password {
		return nil, errors.New("wrong password")
	}
	token, err := jwt.GenerateToken(*user)
	if err != nil {
		return nil, err
	}
	return UserLoginResp{Token: token}, nil
}
