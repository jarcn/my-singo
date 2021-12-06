package api

import (
	"my-singo/conf"
	"my-singo/middleware"
	"my-singo/serializer"

	"github.com/gin-gonic/gin"
)

// 国际化测试
// @Summary 国际化
// @Description 国际化测试
// @Tags 国际化测试接口
// @Success 200
// @Router /lan [get]
func LanTest(c *gin.Context) {
	msg := conf.Message(middleware.Language, "Tag.required")
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  msg,
	})
}
