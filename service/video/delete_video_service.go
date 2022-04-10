package video

import (
	"log"
	"os"
	"path"
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// DeleteVideoService 视频删除服务
type DeleteVideoService struct{}

// 视频删除
func (service *DeleteVideoService) Delete(c *gin.Context) serializer.Response {
	video := model.Video{}
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	vid := c.Param("vid")
	//获取视频信息
	err := model.DB.First(&video, vid).Error
	if err == nil {
		if user.ID != uint(video.Uid) {
			log.Println("uid: ", user.ID, " ", "vuid:", video.Uid)
			return serializer.Response{
				Code: 404,
				Msg:  "用户ID与视频UID不匹配",
			}
		}

		if err := os.Remove("G:/videoResources/" + video.Path); err != nil {
			return serializer.Response{
				Code:  50003,
				Msg:   "视频删除失败",
				Error: err.Error(),
			}
		}
		if err := os.Remove("G:/videoResources/" + video.Cover); err != nil {
			return serializer.Response{
				Code:  50003,
				Msg:   "视频删除失败",
				Error: err.Error(),
			}
		}

		if err := os.Remove(path.Join(DefaultImgPath + video.Cover)); err != nil {
			return serializer.Response{
				Code:  50003,
				Msg:   "视频封面删除失败",
				Error: err.Error(),
			}
		}

		if err = model.DB.Delete(&video).Error; err != nil {
			return serializer.Response{
				Code:  50003,
				Msg:   "视频删除失败",
				Error: err.Error(),
			}
		}

		return serializer.Response{
			Code: 200,
			Msg:  "成功",
		}

	} else {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}
}
