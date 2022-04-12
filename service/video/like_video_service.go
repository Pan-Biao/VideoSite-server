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

	video, re := funcs.GetVideo(vid)
	if re != nil {
		return re.(serializer.Response)
	}

	videoLike := model.VideoLike{
		Uid: user.ID,
		Vid: video.ID,
	}
	temp := model.VideoLike{}
	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("vid = ?", video.ID)
	db.First(&temp)
	if temp.Uid == videoLike.Uid {
		return serializer.Response{
			Code: 200,
			Data: true,
			Msg:  "已经点赞过了",
		}
	}

	if re := funcs.SQLErr(model.DB.Create(&videoLike).Error); re != nil {
		return re.(serializer.Response)
	}

	//播放量加1
	video.LikeNumber = video.LikeNumber + 1
	if re := funcs.SQLErr(model.DB.Save(&video).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "点赞成功",
	}
}
