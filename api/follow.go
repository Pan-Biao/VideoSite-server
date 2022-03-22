package api

import (
	"vodeoWeb/service/follow"

	"github.com/gin-gonic/gin"
)

// CreateFollow 关注
func CreateFollow(c *gin.Context) {
	service := follow.CreateFollowService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// ListFollower 关注列表接口
func ListFollower(c *gin.Context) {
	service := follow.ListFollowerService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// ListFans 粉丝列表接口
func ListFans(c *gin.Context) {
	service := follow.ListFansService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// DeleteFollow 取消关注接口
func DeleteFollow(c *gin.Context) {
	service := follow.DeleteFollowService{}
	res := service.Delete(c)
	c.JSON(200, res)
}
