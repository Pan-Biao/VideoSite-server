package api

import (
	"vodeoWeb/serializer"
	"vodeoWeb/service/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	service := user.UserRegisterService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	service := user.UserLoginService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// UpdateUser 修改用户信息接口
func UpdateUser(c *gin.Context) {
	service := user.UpdateUserService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.Response{
		Code: 200,
		Data: serializer.BuildUser(*user),
		Msg:  "成功"}
	c.JSON(200, res)
}

//获取用户信息
func UserInformation(c *gin.Context) {
	service := user.GetUserInformationService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 200,
		Msg:  "登出成功",
	})
}
