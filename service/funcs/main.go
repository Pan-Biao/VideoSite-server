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
func GetVideo(id string) (model.Video, interface{}) {
	video := model.Video{}

	re := SQLErr(model.DB.First(&video, id).Error)

	if re != nil {
		return video, re.(serializer.Response)
	}

	return video, nil
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
		return FileErr(err)
	}

	err = ioutil.WriteFile(str, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		return FileErr(err)
	}
	return nil
}

// 数据库错误
func SQLErr(err error) interface{} {
	if err != nil {
		return serializer.Response{
			Code:  999999,
			Msg:   "数据库出错",
			Error: err.Error(),
		}
	}
	return nil
}

//文件读取错误
func FileErr(err error) interface{} {
	if err != nil {
		return serializer.Response{
			Code:  999998,
			Msg:   "读取文件出错",
			Error: err.Error(),
		}
	}
	return nil
}

//文件保存错误
func SaveFileErr(err error) interface{} {
	if err != nil {
		return serializer.Response{
			Code:  999997,
			Msg:   "保存文件出错",
			Error: err.Error(),
		}
	}
	return nil
}
