package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

var Language string

// 获取请求语言,全局使用
func LanguageFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if lan := c.GetHeader("Accept-Language"); lan == "" {
			Language = "en"
		} else {
			Language = convertLan(lan)
		}
	}
}

func convertLan(lan string) string {
	str := strings.Split(lan, ",")
	return strings.Split(str[0], "-")[0]
}
