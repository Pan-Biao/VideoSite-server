package user

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

// UpdateUserService 修改用户信息的服务
type UpdateUserService struct {
	Nickname string `form:"nickname" json:"nickname" binding:"min=2,max=8"`
}

const DefaultPath = "G:/videoResources"
const avater = "avater"

// Login 用户登录函数
func (service *UpdateUserService) Update(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}

	//获取上传文件
	file, err := c.FormFile("img")
	if err == nil {
		log.Println("filename:", file.Filename)
		id := strconv.FormatUint(uint64(user.ID), 10)
		//新文件名
		newName := id + util.Intercept(file.Filename)
		log.Println(newName)
		//保存路径 创建文件夹
		dst := path.Join(DefaultPath, id, avater)
		os.MkdirAll(dst, 0777)
		//文件路径
		filePath := path.Join(DefaultPath, id, avater, newName)
		log.Println(filePath)
		//保存文件
		if c.SaveUploadedFile(file, filePath) != nil {
			return serializer.Response{
				Code:  500006,
				Msg:   "读取文件错误",
				Error: err.Error(),
			}
		}
		//更新数据库数据
		user.Avatar = filePath
	}

	//更新数据库数据
	user.Nickname = service.Nickname
	model.DB.Save(&user)

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildUser(user),
		Msg:  "成功",
	}
}
