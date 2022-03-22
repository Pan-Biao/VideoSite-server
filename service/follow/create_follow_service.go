package follow

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

type CreateFollowService struct{}

// CreateFollowService 关注的服务
func (service *CreateFollowService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	fid := c.Param("fid")
	follower := model.User{}
	if err := model.DB.First(&follower, fid).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "无法查找到需要关注的用户",
			Error: err.Error(),
		}
	}

	follow := model.Follow{
		Follower: follower.ID,
		Fans:     user.ID,
	}

	if err := model.DB.Create(&follow).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "关注失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "关注成功",
	}
}
