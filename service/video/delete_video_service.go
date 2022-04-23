package video

import (
	"log"
	"os"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// DeleteVideoService 视频删除服务
type DeleteVideoService struct{}

// 视频删除
func (service *DeleteVideoService) Delete(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	vid := c.Param("vid")
	//获取视频信息
	video := funcs.GetVideo(vid)
	if video == (model.Video{}) {
		return serializer.DBErr("", nil)
	}

	if user.ID != video.Uid {
		log.Println("uid: ", user.ID, " ", "vuid:", video.Uid)
		return serializer.CheckNoRight()
	}

	if err := os.Remove("G:/videoResources/" + video.Path); err != nil {
		video.State = false
		if err := model.DB.Save(&video).Error; err != nil {
			return serializer.DBErr("", err)
		}
		return serializer.FileErr("", err)
	}
	video.Path = ""

	if err := os.Remove("G:/videoResources/" + video.Cover); err != nil {
		video.State = false
		if err := model.DB.Save(&video).Error; err != nil {
			return serializer.DBErr("", err)
		}
		return serializer.FileErr("", err)
	}

	if err := model.DB.Unscoped().Delete(&video).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("视频删除成功", true)

}
