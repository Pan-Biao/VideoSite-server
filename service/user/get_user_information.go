package user

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// UpdateUserService 修改用户信息的服务
type GetUserInformationService struct{}

// Login 用户登录函数
func (service *GetUserInformationService) Get(c *gin.Context) serializer.Response {
	user := model.User{}
	uid := c.Param("uid")
	if err := model.DB.First(&user, uid).Error; err != nil {
		return serializer.Response{
			Code: 404,
			Msg:  err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildUser(user),
		Msg:  "成功",
	}
}
