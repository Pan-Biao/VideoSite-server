package video

import (
	"log"
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// 视频播放服务
type PlayVideoService struct{}

func (service *PlayVideoService) Add(c *gin.Context) {
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

	log.Println(video.PlayNumber)
	video.PlayNumber = video.PlayNumber + 1
	//播放量加1
	if err := model.DB.Save(&video).Error; err != nil {
		c.JSON(500, serializer.Response{
			Code:  50002,
			Msg:   "播放量增加失败",
			Error: err.Error(),
		})
	}

	c.JSON(200, serializer.Response{
		Code: 200,
		Data: true,
	})
}
