package follow

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type DeleteFollowService struct{}

// DeleteFollowService 取消关注服务
func (service *DeleteFollowService) Delete(c *gin.Context) serializer.Response {
	follow := model.Follow{}
	fid := c.Param("fid")
	//获取当前用户
	user := funcs.GetUser(c)
	db := model.DB
	db = db.Where("follower = ? and fans = ?", fid, user.ID)

	if re := funcs.SQLErr(db.First(&follow).Error); re != nil {
		return re.(serializer.Response)
	}

	if re := funcs.SQLErr(model.DB.Unscoped().Delete(&follow).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "成功",
	}
}
