package video

import (
	"log"
	"os"
	"path"
	"strconv"
	"unicode/utf8"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"
	"vodeoWeb/util"

	"github.com/gin-gonic/gin"
)

// CreateVideoService 视频投稿的服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" `
	Info  string `form:"info" json:"info" `
	Said  uint   `form:"said" json:"said" `
}

const DefaultVideoPath = "G:/videoResources/video"
const DefaultImgPath = "G:/videoResources/cover"

// 视频投稿的服务
func (service *CreateVideoService) Create(c *gin.Context) serializer.Response {
	if utf8.RuneCountInString(service.Title) < 1 || utf8.RuneCountInString(service.Title) > 40 {
		return serializer.ParamErr("标题长度应为1-40个字")
	}
	if utf8.RuneCountInString(service.Info) > 300 {
		return serializer.ParamErr("视频简介长度应为300个字以下")
	}
	if service.Said == 0 {
		return serializer.ParamErr("请选择分区")
	}
	count := int64(0)
	model.DB.Model(&model.SubArea{}).First(&model.SubArea{}, service.Said).Count(&count)
	if count == 0 {
		return serializer.ParamErr("分区不存在")
	}

	//获取当前用户
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
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

	if err := model.DB.Create(&video).Error; err != nil {
		return serializer.DBErr("", err)
	}

	id := strconv.FormatUint(uint64(user.ID), 10)
	vid := strconv.FormatUint(uint64(video.ID), 10)
	//上传视频
	{
		videoFile, head, err := c.Request.FormFile("video")

		if err != nil {
			model.DB.Unscoped().Delete(&video)
			return serializer.ParamErr("请上传视频文件")
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
			model.DB.Unscoped().Delete(&video)
			return re.(serializer.Response)
		}

		video.Path = path.Join("video", id, newVideoName)
	}

	//上传图片
	{
		vimgFile, err := c.FormFile("vimg")

		if err != nil {
			model.DB.Unscoped().Delete(&video)
			return serializer.ParamErr("请上传视频封面文件")
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
		if err := c.SaveUploadedFile(vimgFile, imgFilePath); err != nil {
			model.DB.Unscoped().Delete(&video)
			return serializer.FileErr("", err)
		}
		video.Cover = path.Join("cover", id, newVimgName)
	}

	//更新投稿状态
	video.State = true
	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("创建视频成功", serializer.BuildVideo(video))
}
