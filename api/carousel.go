package api

import (
	"vodeoWeb/service/carousel"

	"github.com/gin-gonic/gin"
)

// 添加轮播推荐
func AddCarousel(c *gin.Context) {
	service := carousel.AddCarouselService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Add(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 轮播推荐列表
func ListCarousel(c *gin.Context) {
	service := carousel.ListCarouselService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 删除轮播推荐
func DeleteCarousel(c *gin.Context) {
	service := carousel.DeleteCarouselService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
