package service

import (
	"my-singo/conf"
	"my-singo/middleware"
	"my-singo/model"
	"my-singo/serializer"
	"my-singo/util"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr(conf.Message(middleware.Language, "UserName.Error"), nil)
	}

	if !user.CheckPassword(service.Password) {
		return serializer.ParamErr(conf.Message(middleware.Language, "Password.Error"), nil)
	}

	tokenUser := util.User{
		UserName: service.UserName,
		PassWrod: service.Password,
	}

	//生成token
	token, err := util.GenToken(tokenUser)
	if err != nil {
		return serializer.Err(serializer.CodeSystemErr, conf.Message(middleware.Language, "System.Error"), nil)
	}
	// 设置session
	// service.setSession(c, user)
	return serializer.Response{
		Code: http.StatusOK,
		Data: gin.H{"token": token},
		Msg:  conf.Message(middleware.Language, "Login.Success"),
	}
}
