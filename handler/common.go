package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type ParserHandler interface {
	Handler(c *gin.Context) (interface{}, error)
}

type response struct {
	Status string
	Time   time.Time
	Data   interface{}
}

func handler(c *gin.Context, h interface{}) {
	if handler, ok := h.(ParserHandler); ok {
		err := c.Bind(handler)
		if err != nil {
			fmt.Println("[Error] " + err.Error())
		}
		data, err := handler.Handler(c)
		if err != nil {
			fmt.Println("[Error] " + err.Error())
		}
		doResp(c, data, err)
	}
}

func doResp(c *gin.Context, data interface{}, err error) {
	resp := response{Time: time.Now()}
	if err != nil {
		resp.Status = "Error"
		resp.Data = err.Error()
	} else {
		resp.Status = "Success"
		resp.Data = data
	}
	c.JSON(200, resp)
}

func MainHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "protect"})
}
