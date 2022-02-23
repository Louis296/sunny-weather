package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/service/moment_service"
)

func (r Handler) CreateMoment20220101(c *gin.Context) {
	handler(c, &moment_service.CreateMomentReq{})
}

func (r Handler) DeleteMoment20220101(c *gin.Context) {
	handler(c, &moment_service.DeleteMomentReq{})
}

func (r Handler) UpdateMoment20220101(c *gin.Context) {
	handler(c, &moment_service.UpdateMomentReq{})
}

func (r Handler) ListSelfMoment20220101(c *gin.Context) {
	handler(c, &moment_service.ListSelfMomentReq{})
}

func (r Handler) ListMoment20220101(c *gin.Context) {
	handler(c, &moment_service.ListMomentReq{})
}

func (r Handler) ThumbUpMoment20220101(c *gin.Context) {
	handler(c, &moment_service.ThumbUpMomentReq{})
}
