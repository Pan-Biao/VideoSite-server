package video

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 封禁用户的服务
type VideoUnsealService struct{}

// 用户封禁函数
func (service *VideoUnsealService) Unseal(c *gin.Context) serializer.Response {
	//检测root
	if !funcs.CheckRoot(c) {
		return serializer.CheckNoRight()
	}
	vid := c.Param("vid")
	//获取封禁用户
	video := funcs.GetVideo(vid)
	if video == (model.Video{}) {
		return serializer.DBErr("", nil)
	}

	//更新数据库数据
	video.State = true
	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("封禁成功", true)
}
