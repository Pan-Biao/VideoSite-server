package user

import (
	"unicode/utf8"
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" `
	Password string `form:"password" json:"password" `
	Token    bool   `form:"token" json:"token"`
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

	if utf8.RuneCountInString(service.UserName) < 6 || utf8.RuneCountInString(service.UserName) > 12 {
		return serializer.ParamErr("账号应为6-12位数")
	}
	if utf8.RuneCountInString(service.Password) < 6 || utf8.RuneCountInString(service.Password) > 16 {
		return serializer.ParamErr("密码应为6-16位数")
	}

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误")
	}

	if !user.CheckPassword(service.Password) {
		return serializer.ParamErr("账号或密码错误")
	}

	var token string
	var tokenExpire int64
	var err error
	if service.Token {
		token, tokenExpire, err = user.MakeToken()
		if err != nil {
			return serializer.DBErr("redis err", err)
		}
	} else {
		// web端设置session
		service.setSession(c, user)
	}

	data := serializer.BuildUserToken(user)
	data.Token = token
	data.TokenExpire = tokenExpire

	return serializer.ReturnData("登录成功", data)
}
