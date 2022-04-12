package follow

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// ListFansService 列表
type ListFansService struct{}

// 列表服务
func (service *ListFansService) List(c *gin.Context) serializer.Response {
	follow := []model.Follow{}
	uid := c.Param("uid")

	if re := funcs.SQLErr(model.DB.Where("follower = ?", uid).Find(&follow).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildFans(follow),
		Msg:  "成功",
	}
}
