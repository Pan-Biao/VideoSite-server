package video

import (
	"log"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/util"

	"github.com/gin-gonic/gin"
)

// // PlayVideoService 视频播放服务
type PlayVideoService struct{}

// // 视频更新
// func (service *PlayVideoService) Play(c *gin.Context) {
// 	//查找对应视频
// 	vid := util.Intercept2(c.Param("vid"))
// 	video := model.Video{}
// 	model.DB.First(&video, vid)

// 	videoFile, err := os.Open(video.Path)
// 	if err != nil {
// 		c.JSON(500, serializer.Response{
// 			Code:  50006,
// 			Msg:   "视频播放出错",
// 			Error: err.Error(),
// 		})
// 	}
// 	defer videoFile.Close()
// 	http.ServeContent(c.Writer, c.Request, path, time.Now(), videoFile)
// }

func (service *PlayVideoService) Add(c *gin.Context) {
	//查找对应视频
	vid := util.Intercept2(c.Param("vid"))
	video := model.Video{}
	if err := model.DB.First(&video, vid).Error; err != nil {
		c.JSON(500, serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		})
	} else {
		log.Println(video.PlayNumber)
		video.PlayNumber = video.PlayNumber + 1
		//播放量加1
		if err := model.DB.Save(&video).Error; err != nil {
			c.JSON(500, serializer.Response{
				Code:  50002,
				Msg:   "视频更新失败",
				Error: err.Error(),
			})
		}
	}
}
