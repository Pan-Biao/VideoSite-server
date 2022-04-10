package follow

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// ListFansService 分区列表
type ListFansService struct{}

// List 分区列表服务
func (service *ListFansService) List(c *gin.Context) serializer.Response {
	follow := []model.Follow{}
	uid := c.Param("uid")

	if err := model.DB.Where("follower = ?", uid).Find(&follow).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "查询错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildFollows(follow),
		Msg:  "成功",
	}
}
