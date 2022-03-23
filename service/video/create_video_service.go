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

const DefaultVideoPath = "G:/videoResources/video"
const DefaultImgPath = "G:/videoResources/cover"

// 视频投稿的服务
func (service *CreateVideoService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}

	video := model.Video{
		Title:      service.Title,
		Info:       service.Info,
		Path:       "",
		State:      false,
		PlayNumber: 0,
		Uid:        user.ID,
		Said:       service.Said,
	}

	model.DB.Create(&video)

	//获取上传文件
	videoFile, err := c.FormFile("video")
	if err != nil {
		model.DB.Delete(&video)
		return serializer.Response{
			Code:  50005,
			Msg:   "上传视频失败,可能是文件名太长",
			Error: err.Error(),
		}
	}
	vimgFile, err := c.FormFile("vimg")
	if err == nil {
		log.Println(videoFile.Filename, vimgFile.Filename)
		id := strconv.FormatUint(uint64(user.ID), 10)
		vid := strconv.FormatUint(uint64(video.ID), 10)
		//上传
		newVideoName := vid + util.Intercept(videoFile.Filename)
		newVimgName := vid + util.Intercept(vimgFile.Filename)
		log.Println(newVideoName, newVimgName)
		//保存路径 创建文件夹
		videoDst := path.Join(DefaultVideoPath, id)
		imgDst := path.Join(DefaultImgPath, id)
		os.MkdirAll(videoDst, 0777)
		os.MkdirAll(imgDst, 0777)
		//文件路径
		videoFilePath := path.Join(DefaultVideoPath, id, newVideoName)
		imgFilePath := path.Join(DefaultImgPath, id, newVimgName)
		log.Println(videoFilePath, newVimgName)
		//保存video文件
		if c.SaveUploadedFile(videoFile, videoFilePath) != nil {
			model.DB.Delete(&video)
			return serializer.Response{
				Code:  500006,
				Msg:   "读取文件错误",
				Error: err.Error(),
			}
		}
		//保存img文件
		if c.SaveUploadedFile(vimgFile, imgFilePath) != nil {
			model.DB.Delete(&video)
			return serializer.Response{
				Code:  500006,
				Msg:   "读取文件错误",
				Error: err.Error(),
			}
		}
		//更新投稿状态
		video.State = true
		video.Path = videoFilePath
		video.Cover = path.Join(id, newVimgName)
		model.DB.Save(&video)

		return serializer.Response{
			Code: 200,
			Data: serializer.BuildVideo(video),
			Msg:  "成功",
		}

	} else {
		model.DB.Delete(&video)
		return serializer.Response{
			Code:  50005,
			Msg:   "上传视频封面失败,可能是文件名太长",
			Error: err.Error(),
		}
	}
}
