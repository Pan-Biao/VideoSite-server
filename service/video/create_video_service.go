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

	if re := funcs.SQLErr(model.DB.Create(&video).Error); re != nil {
		return re.(serializer.Response)
	}

	id := strconv.FormatUint(uint64(user.ID), 10)
	vid := strconv.FormatUint(uint64(video.ID), 10)
	//上传视频
	{
		videoFile, head, err := c.Request.FormFile("video")

		if re := funcs.FileErr(err); re != nil {
			model.DB.Delete(&video)
			return re.(serializer.Response)
		}

		//读取
		log.Println(head.Filename)
		newVideoName := vid + util.Intercept(head.Filename)
		log.Println(newVideoName)
		videoDst := path.Join(DefaultVideoPath, id)
		os.MkdirAll(videoDst, 0777)
		videoFilePath := path.Join(DefaultVideoPath, id, newVideoName)
		log.Println(videoFilePath)

		//保存video文件
		if re := funcs.SaveFile(&videoFile, videoFilePath); re != nil {
			return re.(serializer.Response)
		}

		video.Path = path.Join("video", id, newVideoName)
	}

	//上传图片
	{
		vimgFile, err := c.FormFile("vimg")

		if re := funcs.FileErr(err); re != nil {
			model.DB.Delete(&video)
			return re.(serializer.Response)
		}

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
		if re := funcs.SaveFileErr(c.SaveUploadedFile(vimgFile, imgFilePath)); re != nil {
			model.DB.Delete(&video)
			return re.(serializer.Response)
		}
		video.Cover = path.Join("cover", id, newVimgName)
	}

	//更新投稿状态
	video.State = true
	if re := funcs.SQLErr(model.DB.Save(&video).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildVideo(video),
		Msg:  "成功",
	}
}
