package api

import (
	"my-singo/conf"
	"my-singo/middleware"
	"my-singo/serializer"

	"github.com/gin-gonic/gin"
)

func LanTest(c *gin.Context) {
	msg := conf.Message(middleware.Language, "Tag.required")
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  msg,
	})
}
