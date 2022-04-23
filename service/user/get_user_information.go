package user

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// 获取用户信息的服务
type GetUserInformationService struct{}

func (service *GetUserInformationService) Get(c *gin.Context) serializer.Response {
	user := model.User{}
	uid := c.Param("uid")

	if err := model.DB.First(&user, uid).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("成功", serializer.BuildUser(user))
}
