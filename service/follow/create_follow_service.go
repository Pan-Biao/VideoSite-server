package follow

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type CreateFollowService struct{}

// CreateFollowService 关注的服务
func (service *CreateFollowService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	fid := c.Param("fid")

	follower := model.User{}
	count := int64(0)
	model.DB.First(&follower, fid).Count(&count)
	if count == 0 {
		return serializer.ReturnData("关注用户不存在", false)
	}

	follow := model.Follow{
		Follower: follower.ID,
		Fans:     user.ID,
	}
	temp := model.Follow{}
	model.DB.Where("follower = ?", follow.Follower).Where("fans = ?", follow.Fans).First(&temp).Count(&count)
	if count != 0 {
		return serializer.ReturnData("已关注", true)
	}

	if err := model.DB.Create(&follow).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("关注成功", true)
}
