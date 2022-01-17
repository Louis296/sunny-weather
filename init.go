package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/sunny-weather/conf"
	"github.com/louis296/sunny-weather/handler"
)

func Init(r *gin.Engine) {
	// init router
	r.GET("/ping", handler.Ping)

	configure, err := conf.GetConf()
	if err != nil {
		panic(err)
	}
	err = r.Run(":" + configure.Server.Port)
	if err != nil {
		panic(err)
	}
}
