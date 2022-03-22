package video

import (
	"log"
	"net/http"
	"os"
	"time"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/util"

	"github.com/gin-gonic/gin"
)

// PlayVideoService 视频播放服务
type PlayVideoService struct{}

// 视频更新
func (service *PlayVideoService) Play(c *gin.Context) {
	vid := util.Intercept2(c.Param("vid"))
	//查找对应视频
	video := model.Video{}
	log.Println(vid)
	err := model.DB.First(&video, vid).Error
	if err == nil {
		//播放量加1
		video.PlayNumber = video.PlayNumber + 1
		err = model.DB.Save(&video).Error
		if err != nil {
			c.JSON(500, serializer.Response{
				Code:  50002,
				Msg:   "视频更新失败",
				Error: err.Error(),
			})
		}
		//视频服务
		path := video.Path
		log.Println(path)
		videoFile, err := os.Open(path)
		if err != nil {
			c.JSON(500, serializer.Response{
				Code:  50006,
				Msg:   "视频播放出错",
				Error: err.Error(),
			})
		}
		defer videoFile.Close()
		http.ServeContent(c.Writer, c.Request, path, time.Now(), videoFile)

	} else {
		c.JSON(500, serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		})
	}
}
