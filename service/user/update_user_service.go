package user

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

// UpdateUserService 修改用户信息的服务
type UpdateUserService struct {
	Nickname string `form:"nickname" json:"nickname" `
	Info     string `form:"info" json:"info"`
}

const DefaultHeadPortraitPath = "G:/videoResources/head_portrait"

func (service *UpdateUserService) Update(c *gin.Context) serializer.Response {

	if utf8.RuneCountInString(service.Nickname) < 2 && utf8.RuneCountInString(service.Nickname) > 8 {
		return serializer.ParamErr("昵称长度应为2-8个字")
	}
	if utf8.RuneCountInString(service.Info) > 40 {
		return serializer.ParamErr("简介长度应为40个字以下")
	}

	//获取当前用户
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	//获取上传文件
	file, _ := c.FormFile("img")
	if file != nil {
		id := strconv.FormatUint(uint64(user.ID), 10)
		//新文件名
		newName := id + util.Intercept(file.Filename)
		//保存路径 创建文件夹
		dst := path.Join(DefaultHeadPortraitPath, id)
		os.MkdirAll(dst, 0777)
		//文件路径
		filePath := path.Join(DefaultHeadPortraitPath, id, newName)
		//保存文件
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			return serializer.FileErr("", err)
		}
		//更新数据库数据
		user.HeadPortrait = path.Join("head_portrait", id, newName)
	}

	//更新数据库数据
	user.Nickname = service.Nickname
	user.Info = service.Info
	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("成功", serializer.BuildUser(user))
}
