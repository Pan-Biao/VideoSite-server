package video

import (
	"log"
	"os"
	"path"
	"strconv"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/util"

	"github.com/gin-gonic/gin"
)

// CreateVideoService 视频投稿的服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Info  string `form:"info" json:"info" binding:"min=0,max=300"`
	Said  uint   `form:"said" json:"said" binding:"required"`
}

const DefaultPath = "G:/videoResources"
const vpath = "video"

// 视频投稿的服务
func (service *CreateVideoService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}

	log.Println(service.Said)
	// said, err := strconv.Atoi(service.Said)
	// if err != nil {
	// 	return serializer.Response{
	// 		Code: 200,
	// 		Msg:  "said参数错误",
	// 	}
	// }

	video := model.Video{
		Title:      service.Title,
		Info:       service.Info,
		Path:       "",
		State:      false,
		PlayNumber: 0,
		Uid:        user.ID,
		Said:       service.Said,
	}
	lastVideo := model.Video{}
	if err := model.DB.Last(&lastVideo).Error; err == nil {
		return serializer.Response{
			Code: 404,
			Data: serializer.BuildVideo(video),
			Msg:  "数据库查询错误",
		}
	}
	tempvid := lastVideo.ID + 1

	//获取上传文件
	file, err := c.FormFile("video")
	if err == nil {
		log.Println(file.Filename)
		id := strconv.FormatUint(uint64(user.ID), 10)
		vid := strconv.FormatUint(uint64(tempvid), 10)
		//上传
		newName := vid + util.Intercept(file.Filename)
		log.Println(newName)
		//保存路径 创建文件夹
		dst := path.Join(DefaultPath, id, vpath)
		os.MkdirAll(dst, 0777)
		//文件路径
		filePath := path.Join(DefaultPath, id, vpath, newName)
		log.Println(filePath)
		//保存文件
		if c.SaveUploadedFile(file, filePath) != nil {
			return serializer.Response{
				Code:  500006,
				Msg:   "读取文件错误",
				Error: err.Error(),
			}
		}
		//更新投稿状态
		video.State = true
		video.Path = filePath
		model.DB.Save(&video)

		return serializer.Response{
			Code: 200,
			Data: serializer.BuildVideo(video),
			Msg:  "成功",
		}

	} else {
		return serializer.Response{
			Code:  50005,
			Msg:   "上传失败,可能是文件名太长",
			Error: err.Error(),
		}
	}
}
