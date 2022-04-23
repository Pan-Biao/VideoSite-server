package funcs

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"regexp"
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// 验证Root用户
func CheckRoot(c *gin.Context) bool {
	user := GetUser(c)
	return user.Root
}

//获取当前用户
func GetUser(c *gin.Context) model.User {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	return user
}

//获取视频
func GetVideo(id string) model.Video {
	video := model.Video{}
	model.DB.First(&video, id)
	return video
}

//利用正则表达式压缩字符串，去除空格或制表符
func CompressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	strss := "\\s+"
	reg := regexp.MustCompile(strss)
	return reg.ReplaceAllString(str, "")
}

//保存文件
func SaveFile(file *multipart.File, str string) interface{} {
	data, err := ioutil.ReadAll(*file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		return serializer.FileErr("", err)
	}

	err = ioutil.WriteFile(str, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		return serializer.FileErr("", err)
	}
	return nil
}
