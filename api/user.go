package api

import (
	"my-singo/cache"
	"my-singo/serializer"
	"my-singo/service"
	"strings"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
// @Tags 用户接口
// @Summary 用户注册
// @Description 用户注册接口
// @Success 200
// @Router /user/register [post]
// @Param nickname formData string true "昵称"
// @Param user_name formData string true "姓名"
// @Param password formData string true "密码"
// @Param password_confirm formData string true "再次确认密码"
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
// @Tags 用户接口
// @Summary 用户登陆
// @Description 用户登陆接口
// @Success 200
// @Router /user/login [post]
// @Param user_name formData string true "姓名"
// @Param password formData string true "密码"
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
// @Tags 用户接口
// @Summary 用户详情
// @Description 用户详情接口
// @Success 200
// @Router /user/me [get]
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
// @Tags 用户接口
// @Summary 用户登出
// @Description 用户登出接口
// @Success 200
// @Router /user/logout [delete]
func UserLogout(c *gin.Context) {
	// token 失效问题需要解决, 将token 存入 redis
	// 用户退出登陆时将 token 加入黑名单
	token := getToken(c)
	cache.RedisClient.Del(token)
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

func getToken(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	return parts[1]
}
