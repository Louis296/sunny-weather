package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/conf"
	"github.com/louis296/sunny-weather/dao"
	"github.com/louis296/sunny-weather/handler"
	"github.com/louis296/sunny-weather/middleware"
	"github.com/louis296/sunny-weather/pkg/jwt"
	"strconv"
)

func Init(r *gin.Engine) {
	// init router
	r.GET("/ping", handler.Ping)
	r.POST("/register", handler.UserRegister)
	r.POST("/login", handler.UserLogin)
	r.Use(middleware.JWT())
	{
		r.Any("/v1", handler.MainHandler)
	}

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

	// init jwt secret
	jwt.Secret = configure.Jwt.Secret

	err = r.Run(":" + strconv.Itoa(configure.Server.Port))
	if err != nil {
		panic(err)
	}
}
