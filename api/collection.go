package api

import (
	"vodeoWeb/service/collection"

	"github.com/gin-gonic/gin"
)

// CreateCollection 创建收藏
func CreateCollection(c *gin.Context) {
	service := collection.CreateCollectionService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// ListCollection 收藏列表接口
func ListCollection(c *gin.Context) {
	service := collection.ListCollectionService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// DeleteCollection 删除收藏接口
func DeleteCollection(c *gin.Context) {
	service := collection.DeleteCollectionService{}
	res := service.Delete(c)
	c.JSON(200, res)
}
