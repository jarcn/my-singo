package api

import (
	"singo/serializer"
	"singo/service"

	"github.com/gin-contrib/sessions"
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
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
