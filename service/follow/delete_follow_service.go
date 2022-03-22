package follow

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

type DeleteFollowService struct{}

// DeleteFollowService 取消关注服务
func (service *DeleteFollowService) Delete(c *gin.Context) serializer.Response {
	follow := model.Follow{}
	fid := c.Param("fid")
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}

	if err := model.DB.Where("follower = ? and fans = ?", fid, user.ID).First(&follow).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "关注不存在",
			Error: err.Error(),
		}
	}

	if err := model.DB.Delete(&follow).Error; err != nil {
		return serializer.Response{
			Code:  60003,
			Msg:   "取消关注失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "成功",
	}
}
