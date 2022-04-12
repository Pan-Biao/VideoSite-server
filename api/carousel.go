package api

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

type CarouselService struct {
	Title string `form:"title" json:"title"`
	Path  string `form:"path" json:"path"`
}

const DefaultCarouselCoverPath = "G:/videoResources/carousel_cover"

// 添加轮播推荐
func AddCarousel(c *gin.Context) {
	carouselService := CarouselService{}
	if err := c.ShouldBind(&carouselService); err != nil {
		c.JSON(500, ErrorResponse(err))
		return
	}
	log.Println(carouselService)

	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	if !user.Root {
		c.JSON(500, "权限不足")
		return
	}

	carousel := model.Carousel{
		Title: carouselService.Title,
		Path:  carouselService.Path,
	}

	if err := model.DB.Create(&carousel).Error; err != nil {
		c.JSON(500, ErrorResponse(err))
		return
	}
	//获取上传文件
	file, err := c.FormFile("file")
	if err == nil {
		log.Println("filename:", file.Filename)
		cid := strconv.FormatUint(uint64(carousel.ID), 10)
		//新文件名
		newName := cid + util.Intercept(file.Filename)
		log.Println(newName)
		//保存路径 创建文件夹
		dst := path.Join(DefaultCarouselCoverPath)
		os.MkdirAll(dst, 0777)
		//文件路径
		filePath := path.Join(DefaultCarouselCoverPath, newName)
		log.Println(filePath)
		//保存文件
		if c.SaveUploadedFile(file, filePath) != nil {
			c.JSON(500, "文件保存失败")
			return
		}
		//更新数据库数据
		carousel.Cover = path.Join("carousel_cover", newName)
	}
	if err := model.DB.Save(&carousel).Error; err != nil {
		c.JSON(500, ErrorResponse(err))
		return
	}
	c.JSON(200, serializer.Response{
		Code: 200,
		Msg:  "成功",
	})
}

type Cb struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
	Cover string `json:"cover"`
}

// 轮播推荐列表
func ListCarousel(c *gin.Context) {
	carousels := []model.Carousel{}

	if err := model.DB.Find(&carousels).Error; err != nil {
		c.JSON(500, "找不到对应数据")
		return
	}
	cb := []Cb{}

	for _, item := range carousels {
		cb = append(cb, Cb{
			ID:    item.ID,
			Title: item.Title,
			Path:  item.Path,
			Cover: item.Cover,
		})
	}

	c.JSON(200, serializer.Response{
		Code: 200,
		Data: cb,
		Msg:  "成功",
	})
}

// 删除轮播推荐
func DeleteCarousel(c *gin.Context) {
	carousel := model.Carousel{}
	cid := c.Param("cid")
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	if !user.Root {
		c.JSON(200, serializer.Response{
			Code: 500,
			Msg:  "权限不足",
		})
		return
	}
	if err := model.DB.First(&carousel, cid).Error; err != nil {
		c.JSON(200, serializer.Response{
			Code: 500,
			Msg:  "找不到对应数据",
		})
		log.Println(carousel)
		return
	}
	errStr := ""
	if err := os.Remove("G:/videoResources/" + carousel.Cover); err != nil {
		errStr = "文件删除失败,"
	}

	if err := model.DB.Unscoped().Delete(&carousel).Error; err != nil {
		c.JSON(200, serializer.Response{
			Code: 500,
			Msg:  errStr + "数据库删除失败",
		})
		return
	}

	c.JSON(200, serializer.Response{
		Code: 200,
		Msg:  errStr + "成功",
	})
}
