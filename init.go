package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/conf"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/handler"
	"strconv"
)

func Init(r *gin.Engine) {
	// init router
	r.GET("/ping", handler.Ping)
	r.POST("/register", handler.UserRegister)
	r.POST("/v1", handler.MainHandler)

	// init configure
	configure, err := conf.GetConf()
	if err != nil {
		panic(err)
	}

	// init database
	err = dao.InitDB(
		configure.Database.URL,
		configure.Database.DatabaseName,
		configure.Database.UserName,
		configure.Database.Password,
		configure.Database.MaxConn,
		configure.Database.MaxOpen,
	)
	if err != nil {
		panic(err)
	}
	err = r.Run(":" + strconv.Itoa(configure.Server.Port))
	if err != nil {
		panic(err)
	}
}
