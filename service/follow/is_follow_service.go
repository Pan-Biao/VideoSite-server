package follow

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// ListFansService 分区列表
type IsFollowService struct{}

// 分区列表服务
func (service *IsFollowService) Is(c *gin.Context) serializer.Response {
	follow := model.Follow{}
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}
	uid := c.Param("uid")

	db := model.DB
	db = db.Where("follower = ?", uid)
	db = db.Where("fans = ?", user.ID)
	count := int64(0)
	db.First(&follow).Count(&count)
	if count == 0 {
		return serializer.ReturnData("未关注", false)
	}

	return serializer.ReturnData("已关注", true)

}
