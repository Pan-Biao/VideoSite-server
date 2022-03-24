package video

import (
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/util"

	"github.com/gin-gonic/gin"
)

// UpdateVideoService 视频更新服务
type UpdateVideoService struct {
	Title *string `form:"title" json:"title" binding:"min=1,max=30"`
	Info  *string `form:"info" json:"info" binding:"min=0,max=300"`
	Said  uint    `form:"said" json:"said" `
	State *bool   `form:"state" json:"state" `
}

// 视频更新
func (service *UpdateVideoService) Update(c *gin.Context) serializer.Response {
	if compressStr(*service.Title) == "" {
		return serializer.Response{
			Code: 404,
			Msg:  "title不能为空",
		}
	}

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

		//更新视频信息
		if service.State != nil {
			video.State = *service.State
		}
		if service.Said != 0 {
			video.Said = service.Said
		}
		if service.Title != nil {
			video.Title = *service.Title
		}
		if service.Info != nil {
			video.Info = *service.Info
		}

		id := strconv.FormatUint(uint64(user.ID), 10)
		vid := strconv.FormatUint(uint64(video.ID), 10)
		//更新视频
		//获取上传文件
		videoFile, _ := c.FormFile("video")
		if videoFile != nil {
			log.Println(videoFile.Filename)
			//上传
			newVideoName := vid + util.Intercept(videoFile.Filename)
			//保存路径 创建文件夹
			videoDst := path.Join(DefaultVideoPath, id)
			os.MkdirAll(videoDst, 0777)
			//文件路径
			videoFilePath := path.Join(DefaultVideoPath, id, newVideoName)
			//保存video文件
			if c.SaveUploadedFile(videoFile, videoFilePath) != nil {
				return serializer.Response{
					Code:  500006,
					Msg:   "读取文件错误",
					Error: err.Error(),
				}
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
			if c.SaveUploadedFile(vimgFile, imgFilePath) != nil {
				return serializer.Response{
					Code:  500006,
					Msg:   "读取文件错误",
					Error: err.Error(),
				}
			}
			video.Cover = path.Join("cover", id, newVimgName)
		}

		if err := model.DB.Save(&video).Error; err != nil {
			return serializer.Response{
				Code:  50002,
				Msg:   "视频更新失败",
				Error: err.Error(),
			}
		}

		return serializer.Response{
			Code: 200,
			Data: serializer.BuildVideo(video),
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

//利用正则表达式压缩字符串，去除空格或制表符
func compressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	strss := "\\s+"
	reg := regexp.MustCompile(strss)
	return reg.ReplaceAllString(str, "")
}
