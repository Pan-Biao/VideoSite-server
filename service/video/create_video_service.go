package video

import (
	"log"
	"os"
	"path"
	"strconv"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"
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
	log.Println("--------------------11111111111----------------------------------------")
	funcs.SQLErr(c, model.DB.Create(&video).Error)
	log.Println("----------------------222222222222--------------------------------------")
	id := strconv.FormatUint(uint64(user.ID), 10)
	vid := strconv.FormatUint(uint64(video.ID), 10)
	//上传视频
	{
		log.Println("-------------------33333333333333333-----------------------------------------")
		videoFile, head, err := c.Request.FormFile("video")
		funcs.FileErr(c, err, func() {
			funcs.SQLErr(c, model.DB.Delete(&video).Error)
		})
		log.Println("-------------------444444444444444444-----------------------------------------")
		//读取
		log.Println(head.Filename)
		newVideoName := vid + util.Intercept(head.Filename)
		log.Println(newVideoName)
		videoDst := path.Join(DefaultVideoPath, id)
		os.MkdirAll(videoDst, 0777)
		videoFilePath := path.Join(DefaultVideoPath, id, newVideoName)
		log.Println(videoFilePath)

		//保存video文件
		funcs.SaveFile(c, &videoFile, videoFilePath)
		// funcs.SaveFileErr(c, c.SaveUploadedFile(videoFile, videoFilePath), func() {
		// 	funcs.SQLErr(c, model.DB.Delete(&video).Error)
		// })
		video.Path = path.Join("video", id, newVideoName)
	}

	//上传图片
	{
		vimgFile, err := c.FormFile("vimg")
		funcs.FileErr(c, err, func() {
			funcs.SQLErr(c, model.DB.Delete(&video).Error)
		})
		log.Println(vimgFile.Filename)
		//上传
		newVimgName := vid + util.Intercept(vimgFile.Filename)
		log.Println(newVimgName)
		//保存路径 创建文件夹
		imgDst := path.Join(DefaultImgPath, id)
		os.MkdirAll(imgDst, 0777)
		//文件路径
		imgFilePath := path.Join(DefaultImgPath, id, newVimgName)
		log.Println(imgFilePath)
		//保存img文件
		if c.SaveUploadedFile(vimgFile, imgFilePath) != nil {
			model.DB.Delete(&video)
			return serializer.Response{
				Code:  500006,
				Msg:   "读取文件错误",
				Error: err.Error(),
			}
		}
		video.Cover = path.Join("cover", id, newVimgName)
	}

	//更新投稿状态
	video.State = true
	funcs.SQLErr(c, model.DB.Save(&video).Error)

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildVideo(video),
		Msg:  "成功",
	}
}
