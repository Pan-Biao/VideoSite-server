package video

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// 视频播放服务
type UnLikeVideoService struct{}

func (service *UnLikeVideoService) UnLike(c *gin.Context) {
	//查找对应视频
	vid := c.Param("vid")
	video := model.Video{}

	if err := model.DB.First(&video, vid).Error; err != nil {
		c.JSON(500, serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		})
	}

	video.LikeNumber = video.LikeNumber - 1
	//播放量加1
	if err := model.DB.Save(&video).Error; err != nil {
		c.JSON(500, serializer.Response{
			Code:  50002,
			Msg:   "播放量增加失败",
			Error: err.Error(),
		})
	}

}
