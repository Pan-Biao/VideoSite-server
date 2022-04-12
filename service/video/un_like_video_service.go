package video

import (
	"log"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 视频播放服务
type UnLikeVideoService struct{}

func (service *UnLikeVideoService) UnLike(c *gin.Context) serializer.Response {
	//查找对应视频
	vid := c.Param("vid")
	user := funcs.GetUser(c)

	video, re := funcs.GetVideo(vid)
	if re != nil {
		return re.(serializer.Response)
	}

	videoLike := model.VideoLike{}
	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("vid = ?", video.ID)

	if re := funcs.SQLErr(db.First(&videoLike).Error); re != nil {
		return serializer.Response{
			Code: 200,
			Data: true,
			Msg:  "已取消点赞",
		}
	}
	log.Println("--------------------------", videoLike)
	if re := funcs.SQLErr(model.DB.Unscoped().Delete(&videoLike).Error); re != nil {
		return re.(serializer.Response)
	}

	//点赞量-1
	video.LikeNumber = video.LikeNumber - 1
	if re := funcs.SQLErr(model.DB.Save(&video).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "取消点赞成功",
	}
}
