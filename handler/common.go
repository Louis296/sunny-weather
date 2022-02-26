package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
)

type Handler struct {
}

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
		resp.Data = struct {
			Message string
		}{Message: err.Error()}
	} else {
		resp.Status = "Success"
		resp.Data = data
	}
	c.JSON(200, resp)
}

func MainHandler(c *gin.Context) {
	action, ok := c.GetQuery("Action")
	if !ok {
		doResp(c, nil, errors.New("no action"))
		c.Abort()
	}
	version, ok := c.GetQuery("Version")
	if !ok {
		doResp(c, nil, errors.New("no version"))
		c.Abort()
	}
	h := Handler{}
	hv := reflect.ValueOf(h)
	f := hv.MethodByName(action + version)
	if !f.IsValid() {
		doResp(c, nil, errors.New("no such api"))
		c.Abort()
	}
	f.Call([]reflect.Value{reflect.ValueOf(c)})
}
