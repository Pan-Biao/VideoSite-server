package api

import (
	"vodeoWeb/service/favorite"

	"github.com/gin-gonic/gin"
)

// CreateFavorite 创建收藏夹
func CreateFavorite(c *gin.Context) {
	service := favorite.CreateFavoriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// ListFavorite 收藏夹列表接口
func ListFavorite(c *gin.Context) {
	service := favorite.ListFavoriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// UpdateFavorite 收藏夹更新接口
func UpdateFavorite(c *gin.Context) {
	service := favorite.UpdateFavoriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// DeleteFavorite 删除收藏夹接口
func DeleteFavorite(c *gin.Context) {
	service := favorite.DeleteFavoriteService{}
	res := service.Delete(c)
	c.JSON(200, res)
}
