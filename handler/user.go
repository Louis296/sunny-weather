package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/service"
)

func UserRegister(c *gin.Context) {
	handler(c, &service.UserRegisterReq{})
}
