package common

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/dao/model"
	"github.com/louis296/sunny-weather/pkg/enum"
)

func GetMD5WithSalt(s string) string {
	salt := "password salt with sunny weather"
	h := md5.New()
	h.Write([]byte(s + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func GetCurrentUser(c *gin.Context) (*model.User, error) {
	user, ok := c.Get(enum.CurrentUser)
	if !ok {
		return nil, errors.New("no current user")
	}
	currentUser, _ := user.(*model.User)
	return currentUser, nil
}
