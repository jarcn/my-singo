package api

import (
	"my-singo/conf"
	"my-singo/serializer"
	"strings"

	"github.com/gin-gonic/gin"
)

// 国际化测试
// @Summary 国际化
// @Description 国际化测试
// @Tags 国际化测试接口
// @Success 200
// @Router /lan [get]
func LanTest(c *gin.Context) {
	recover()
	lan := convertLan(c.GetHeader("Accept-Language"))
	msg := conf.Message(lan, "Tag.required")
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  msg,
	})
}

func convertLan(lan string) string {
	str := strings.Split(lan, ",")
	return strings.Split(str[0], "-")[0]

}
