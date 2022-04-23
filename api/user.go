package api

import (
	"vodeoWeb/cache"
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
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	service := user.UserLoginService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserTokenRefresh 用户刷新token接口
func UserTokenRefresh(c *gin.Context) {
	currUser := CurrentUser(c)
	service := user.UserTokenRefreshService{}
	res := service.Refresh(c, currUser)
	c.JSON(200, res)
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUser(*user)
	c.JSON(200, serializer.Response{
		Code: 200,
		Data: res,
	})
}

// UpdateUser 修改用户信息接口
func UpdateUser(c *gin.Context) {
	service := user.UpdateUserService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//获取用户信息
func UserInformation(c *gin.Context) {
	service := user.GetUserInformationService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Get(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//用户列表接口
func UserList(c *gin.Context) {
	service := user.GetUserListService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//用户封禁
func UserSuspend(c *gin.Context) {
	service := user.UserSuspendService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Suspend(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//用户解封
func UserUnseal(c *gin.Context) {
	service := user.UserUnsealService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Unseal(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	// 移动端登出
	token := c.GetHeader("Authorization")
	if token != "" {
		_ = cache.DelUserToken(token)
	} else {
		// web端登出
		s := sessions.Default(c)
		s.Clear()
		s.Save()
	}
	c.JSON(200, serializer.Response{
		Code: 200,
		Msg:  "登出成功",
		Data: true,
	})
}
