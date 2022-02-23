package moment_service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/pkg/common"
	"strconv"
)

type DeleteMomentReq struct {
	MomentId int
}

func (r *DeleteMomentReq) Handler(c *gin.Context) (interface{}, error) {
	tx := dao.DB.Begin()
	defer func() {
		if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if moment := dao.GetMomentById(r.MomentId); moment == nil {
		return nil, errors.New("no moment with id " + strconv.Itoa(r.MomentId))
	} else {
		currentUser, _ := common.GetCurrentUser(c)
		if currentUser.Id != moment.UserId {
			return nil, errors.New("can't delete moment created by others")
		}
	}
	if err := dao.DeleteMomentById(r.MomentId, tx); err != nil {
		return nil, err
	}
	return nil, nil
}
