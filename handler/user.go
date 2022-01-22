package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/service/user_service"
)

func UserRegister(c *gin.Context) {
	handler(c, &user_service.UserRegisterReq{})
}

func UserLogin(c *gin.Context) {
	handler(c, &user_service.UserLoginReq{})
}

func (r Handler) GetSelfInformation20220101(c *gin.Context) {
	handler(c, &user_service.GetSelfInformationReq{})
}
