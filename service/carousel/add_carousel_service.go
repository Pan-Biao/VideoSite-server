package carousel

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

type AddCarouselService struct {
	Title string `form:"title" json:"title"`
	Path  string `form:"path" json:"path"`
}

const DefaultCarouselCoverPath = "G:/videoResources/carousel_cover"

func (service *AddCarouselService) Add(c *gin.Context) serializer.Response {
	if utf8.RuneCountInString(service.Title) < 1 && utf8.RuneCountInString(service.Title) > 20 {
		return serializer.ParamErr("标题长度应为1-20")
	}
	if !funcs.CheckRoot(c) {
		return serializer.CheckNoRight()
	}

	carousel := model.Carousel{
		Title: service.Title,
		Path:  service.Path,
	}

	if err := model.DB.Create(&carousel).Error; err != nil {
		return serializer.DBErr("", err)
	}
	//获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		return serializer.FileErr("请上传文件", err)
	}

	cid := strconv.FormatUint(uint64(carousel.ID), 10)
	//新文件名
	newName := cid + util.Intercept(file.Filename)
	//保存路径 创建文件夹
	dst := path.Join(DefaultCarouselCoverPath)
	os.MkdirAll(dst, 0777)
	//文件路径
	filePath := path.Join(DefaultCarouselCoverPath, newName)
	//保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return serializer.FileErr("", err)
	}
	//更新数据库数据
	carousel.Cover = path.Join("carousel_cover", newName)

	if err := model.DB.Save(&carousel).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("成功", true)
}
