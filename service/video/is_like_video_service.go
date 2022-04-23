package video

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 视频点赞服务
type IsLikeVideoService struct{}

func (service *IsLikeVideoService) Is(c *gin.Context) serializer.Response {
	//查找对应视频
	vid := c.Param("vid")
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	videoLike := model.VideoLike{}

	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("vid = ?", vid)
	count := int64(0)
	db.First(&videoLike).Count(&count)
	if count == 0 {
		return serializer.ReturnData("未点赞", false)
	}

	return serializer.ReturnData("点赞成功", true)
}
