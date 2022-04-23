package video

import (
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
	count := int64(0)

	model.DB.First(&video, vid).Count(&count)
	if count == 0 {
		c.JSON(200, serializer.ReturnData("没找到视频", false))
	}

	video.PlayNumber = video.PlayNumber + 1
	//播放量加1
	if err := model.DB.Save(&video).Error; err != nil {
		c.JSON(200, serializer.DBErr("", err))
	}

	c.JSON(200, serializer.ReturnData("", true))
}
