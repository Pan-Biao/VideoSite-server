package video

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 取消点赞服务
type UnLikeVideoService struct{}

func (service *UnLikeVideoService) UnLike(c *gin.Context) serializer.Response {
	//查找对应视频
	vid := c.Param("vid")
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	video := funcs.GetVideo(vid)
	if video == (model.Video{}) {
		return serializer.DBErr("", nil)
	}

	videoLike := model.VideoLike{}
	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("vid = ?", video.ID)
	count := int64(0)
	db.First(&videoLike).Count(&count)
	if count == 0 {
		return serializer.ReturnData("未点赞", true)
	}

	if err := model.DB.Unscoped().Delete(&videoLike).Error; err != nil {
		return serializer.DBErr("", err)
	}

	//点赞量-1
	video.LikeNumber = video.LikeNumber - 1
	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("取消点赞成功", true)
}
