package api

import (
	subArea "vodeoWeb/service/sub-area"

	"github.com/gin-gonic/gin"
)

// CreateSubArea 创建分区
func CreateSubArea(c *gin.Context) {
	service := subArea.CreateSubAreaService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// ListVideo 分区列表接口
func ListSubArea(c *gin.Context) {
	service := subArea.ListSubAreaService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// UpdateSubArea 分区更新接口
func UpdateSubArea(c *gin.Context) {
	service := subArea.UpdateSubAreaService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// DeleteSubArea 分区删除接口
func DeleteSubArea(c *gin.Context) {
	service := subArea.DeleteSubAreaService{}
	res := service.Delete(c)
	c.JSON(200, res)
}
