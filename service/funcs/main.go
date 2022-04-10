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
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	return user.Root
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
func SaveFile(c *gin.Context, file *multipart.File, str string) {
	data, err := ioutil.ReadAll(*file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		FileErr(c, err)
		return
	}

	err = ioutil.WriteFile(str, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		FileErr(c, err)
		return
	}
}

// 数据库错误处理
// f 返回前做什么事情
func SQLErr(c *gin.Context, err error, f ...func()) {
	if err != nil {
		for _, v := range f {
			v()
		}
		log.Printf("SQL error: %v", err)
		c.JSON(200, serializer.Response{
			Code:  999999,
			Msg:   "数据库出错",
			Error: err.Error(),
		})
	}
}

//文件读取错误处理
// f 返回前做什么事情
func FileErr(c *gin.Context, err error, f ...func()) {
	if err != nil {
		for _, v := range f {
			v()
		}
		log.Printf("File error: %v", err)
		c.JSON(200, serializer.Response{
			Code:  999998,
			Msg:   "读取文件出错",
			Error: err.Error(),
		})
	}
}

//文件保存错误处理
// f 返回前做什么事情
func SaveFileErr(c *gin.Context, err error, f ...func()) {
	if err != nil {
		for _, v := range f {
			v()
		}
		log.Printf("SaveFile error: %v", err)
		c.JSON(200, serializer.Response{
			Code:  999997,
			Msg:   "保存文件出错",
			Error: err.Error(),
		})
	}
}
