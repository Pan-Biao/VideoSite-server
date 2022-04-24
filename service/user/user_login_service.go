package user

import (
	"fmt"
	"regexp"
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

	if re, _ := regexp.MatchString("^[a-z0-9_-]{6,12}$", service.Password); !re {
		return serializer.ParamErr("用户名格式错误,应为6-12位数字或小写字母")
	}

	if re, _ := regexp.MatchString("^[a-zA-Z0-9]{6,16}$", service.Password); !re {
		return serializer.ParamErr("密码格式错误,应为6-16位数字或大小写字母")
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

	data := serializer.BuildUserToken(user)

	if service.Token {
		token, tokenExpire, err = user.MakeToken()
		if err != nil {
			return serializer.DBErr("redis err", err)
		}
		data.Token = token
		data.TokenExpire = tokenExpire
	} else {
		// web端设置session
		fmt.Println("----------------------", user.ID)
		service.setSession(c, user)
	}

	return serializer.ReturnData("登录成功", data)
}
