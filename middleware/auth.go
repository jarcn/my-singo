package middleware

import (
	"my-singo/serializer"
	"my-singo/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusOK, serializer.CheckLogin())
			ctx.Abort() //结束后续操作
			return
		}
		util.Log().Info("token:", authHeader)

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, serializer.CheckLogin())
			ctx.Abort()
			return
		}

		//解析token包含的信息
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusOK, serializer.CheckLogin())
			ctx.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		ctx.Set("claims", claims)
		ctx.Next() // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息

	}
}
