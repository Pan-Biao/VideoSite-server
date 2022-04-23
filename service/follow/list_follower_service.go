package follow

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// ListFollowerService 分区列表
type ListFollowerService struct{}

// List 分区列表服务
func (service *ListFollowerService) List(c *gin.Context) serializer.Response {
	follow := []model.Follow{}
	uid := c.Param("uid")

	if err := model.DB.Where("fans = ?", uid).Find(&follow).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("成功", serializer.BuildFollowers(follow))
}
