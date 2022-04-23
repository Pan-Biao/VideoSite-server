package follow

import (
	"log"
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
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	db := model.DB
	db = db.Where("follower = ? and fans = ?", fid, user.ID)
	count := int64(0)
	db = db.First(&follow).Count(&count)
	log.Println(count)
	if count == 0 {
		return serializer.ReturnData("已取消关注", true)
	}

	if err := model.DB.Unscoped().Delete(&follow).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("取消关注成功", true)
}
