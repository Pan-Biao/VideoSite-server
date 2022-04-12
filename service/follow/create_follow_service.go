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

	fid := c.Param("fid")

	follower := model.User{}
	if re := funcs.SQLErr(model.DB.First(&follower, fid).Error); re != nil {
		return re.(serializer.Response)
	}

	follow := model.Follow{
		Follower: follower.ID,
		Fans:     user.ID,
	}
	temp := model.Follow{}

	model.DB.Where("follower = ?", follow.Follower).Where("fans = ?", follow.Fans).First(&temp)

	if temp.Fans == follow.Fans {
		return serializer.Response{
			Code: 50001,
			Msg:  "已关注",
		}
	}

	if re := funcs.SQLErr(model.DB.Create(&follow).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "关注成功",
	}
}
