package video

import (
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

// UpdateVideoService 视频更新服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title"`
	Info  string `form:"info" json:"info"`
	Said  uint   `form:"said" json:"said" `
}

// 视频更新
func (service *UpdateVideoService) Update(c *gin.Context) serializer.Response {
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
	model.DB.First(&model.SubArea{}, service.Said).Count(&count)
	if count == 0 {
		return serializer.ParamErr("分区不存在")
	}

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

	if user.ID != uint(video.Uid) {
		return serializer.CheckNoRight()
	}

	//更新视频信息
	video.Said = service.Said
	video.Title = service.Title
	video.Info = service.Info

	id := strconv.FormatUint(uint64(user.ID), 10)

	//更新视频
	//获取上传文件
	videoFile, _ := c.FormFile("video")
	if videoFile != nil {
		//上传
		newVideoName := vid + util.Intercept(videoFile.Filename)
		//保存路径 创建文件夹
		videoDst := path.Join(DefaultVideoPath, id)
		os.MkdirAll(videoDst, 0777)
		//文件路径
		videoFilePath := path.Join(DefaultVideoPath, id, newVideoName)
		//保存video文件
		if err := c.SaveUploadedFile(videoFile, videoFilePath); err != nil {
			return serializer.FileErr("", err)
		}
		//更新投稿状态
		video.Path = path.Join("video", id, newVideoName)
	}

	//更新封面
	vimgFile, _ := c.FormFile("vimg")
	if vimgFile != nil {
		newVimgName := vid + util.Intercept(vimgFile.Filename)
		imgDst := path.Join(DefaultImgPath, id)
		os.MkdirAll(imgDst, 0777)
		imgFilePath := path.Join(DefaultImgPath, id, newVimgName)
		//保存img文件
		if err := c.SaveUploadedFile(vimgFile, imgFilePath); err != nil {
			return serializer.FileErr("", err)
		}
		video.Cover = path.Join("cover", id, newVimgName)

	}

	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("更新成功", serializer.BuildVideo(video))
}
