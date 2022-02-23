package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/service/comment_service"
)

func (r Handler) CreateComment20220101(c *gin.Context) {
	handler(c, &comment_service.CreateCommentReq{})
}

func (r Handler) DeleteComment20220101(c *gin.Context) {
	handler(c, &comment_service.DeleteCommentReq{})
}

func (r Handler) ListComment20220101(c *gin.Context) {
	handler(c, &comment_service.ListCommentReq{})
}
