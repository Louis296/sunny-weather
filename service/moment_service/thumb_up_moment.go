package moment_service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/dao/model"
	"github.com/louis296/sunny-weather/pkg/common"
)

type ThumbUpMomentReq struct {
	MomentId int
}

func (r *ThumbUpMomentReq) Handler(c *gin.Context) (interface{}, error) {
	tx := dao.DB.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	currentUser, _ := common.GetCurrentUser(c)
	if thumbUp := dao.GetMomentUserThumb(r.MomentId, currentUser.Id); thumbUp != nil {
		return nil, errors.New("already liked")
	}
	thumbUp := &model.ThumbUp{
		MomentId: r.MomentId,
		UserId:   currentUser.Id,
	}
	if err := dao.CreateMomentThumbUp(thumbUp, tx); err != nil {
		return nil, err
	}
	moment := dao.GetMomentById(r.MomentId)
	moment.Like++
	if err := dao.SaveMoment(moment, tx); err != nil {
		return nil, err
	}
	return nil, nil
}
