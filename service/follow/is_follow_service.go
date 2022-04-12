package follow

import (
	"log"
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// ListFansService 分区列表
type IsFollowService struct{}

// 分区列表服务
func (service *IsFollowService) Is(c *gin.Context) serializer.Response {
	follow := model.Follow{}
	uid := c.Param("uid")

	if re := model.DB.Where("follower = ?", uid).First(&follow).Error; re != nil {
		return serializer.Response{
			Code: 200,
			Data: false,
			Msg:  "未关注",
		}
	}

	log.Println(follow)

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "已关注",
	}

}
