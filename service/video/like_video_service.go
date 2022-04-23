package video

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 视频点赞服务
type LikeVideoService struct{}

func (service *LikeVideoService) Like(c *gin.Context) serializer.Response {
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

	videoLike := model.VideoLike{
		Uid: user.ID,
		Vid: video.ID,
	}
	count := int64(0)
	temp := model.VideoLike{}
	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("vid = ?", video.ID)
	db.First(&temp).Count(&count)
	if count > 0 {
		serializer.ReturnData("已经点过赞了", true)
	}

	if err := model.DB.Create(&videoLike).Error; err != nil {
		return serializer.DBErr("", err)
	}

	//播放量加1
	video.LikeNumber = video.LikeNumber + 1
	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("点赞成功", true)
}
