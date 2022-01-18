package service

import (
	"github.com/gin-gonic/gin"
)

type UserRegisterReq struct {
	Email    string
	Name     string
	Password string
}

type UserRegisterResp struct {
	Name string
}

func (r *UserRegisterReq) Handler(c *gin.Context) (interface{}, error) {
	return UserRegisterResp{Name: "test"}, nil
}
