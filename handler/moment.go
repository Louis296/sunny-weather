package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/service/moment_service"
)

func (r Handler) CreateMoment20220101(c *gin.Context) {
	handler(c, &moment_service.CreateMomentReq{})
}
