package carousel

import (
	"os"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type DeleteCarouselService struct{}

func (service *DeleteCarouselService) Delete(c *gin.Context) serializer.Response {
	carousel := model.Carousel{}
	cid := c.Param("cid")

	if !funcs.CheckRoot(c) {
		return serializer.CheckNoRight()
	}
	count := int64(0)
	model.DB.First(&carousel, cid).Count(&count)
	if count == 0 {
		return serializer.ReturnData("轮播图不存在", false)
	}

	if err := os.Remove("G:/videoResources/" + carousel.Cover); err != nil {
		return serializer.FileErr("", err)
	}

	if err := model.DB.Unscoped().Delete(&carousel).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("删除成功", true)
}
