package video

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 封禁用户的服务
type VideoSuspendService struct{}

// 用户封禁函数
func (service *VideoSuspendService) Suspend(c *gin.Context) serializer.Response {
	//检测root
	if !funcs.CheckRoot(c) {
		return serializer.Response{
			Code: 500,
			Msg:  "无权限",
		}
	}
	vid := c.Param("vid")
	//获取封禁视频
	video := model.Video{}
	if err := model.DB.First(&video, vid).Error; err != nil {
		return serializer.Response{
			Code: 404,
			Msg:  "查询出错",
		}
	}
	//更新数据库数据
	video.State = false
	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.Response{
			Code: 404,
			Msg:  "保存信息错误",
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "封禁成功",
	}
}
