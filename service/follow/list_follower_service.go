package follow

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// ListFollowerService 分区列表
type ListFollowerService struct{}

// List 分区列表服务
func (service *ListFollowerService) List(c *gin.Context) serializer.Response {
	follow := []model.Follow{}
	uid := c.Param("uid")

	if re := funcs.SQLErr(model.DB.Where("fans = ?", uid).Find(&follow).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildFollowers(follow),
		Msg:  "成功",
	}
}
