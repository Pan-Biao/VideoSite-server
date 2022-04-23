package user

import (
	"vodeoWeb/cache"
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// UserTokenRefreshService 用户刷新token的服务
type UserTokenRefreshService struct {
}

// Refresh 刷新token
func (service *UserTokenRefreshService) Refresh(c *gin.Context, user *model.User) serializer.Response {
	//删除原来的Token
	oldToken := c.GetHeader("Authorization")
	cache.DelUserToken(oldToken)
	//生成新的Token
	newToken, tokenExpire, err := user.MakeToken()
	if err != nil {
		return serializer.DBErr("redis err", err)
	}
	data := serializer.BuildUserToken(*user)
	data.Token = newToken
	data.TokenExpire = tokenExpire
	return serializer.Response{
		Code: 200,
		Data: data,
	}
}
