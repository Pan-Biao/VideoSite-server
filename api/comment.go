package api

import (
	"vodeoWeb/service/comment"

	"github.com/gin-gonic/gin"
)

// 创建评论
func CreateComment(c *gin.Context) {
	service := comment.CreateCommentService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// 评论列表
func ListComment(c *gin.Context) {
	service := comment.ListCommentService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// 删除评论
func DeleteComment(c *gin.Context) {
	service := comment.DeleteCommentService{}
	res := service.Delete(c)
	c.JSON(200, res)
}
