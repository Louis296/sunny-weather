package user_service

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao/model"
	"github.com/louis296/sunny-weather/pkg/enum"
)

type GetSelfInformationReq struct {
	Test string
}

type GetSelfInformationResp struct {
	Name  string
	Email string
}

func (g *GetSelfInformationReq) Handler(c *gin.Context) (interface{}, error) {
	elem, _ := c.Get(enum.CurrentUser)
	user, _ := elem.(*model.User)
	return GetSelfInformationResp{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
