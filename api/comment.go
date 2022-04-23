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
		c.JSON(200, ErrorResponse(err))
	}
}

// 评论列表
func ListComment(c *gin.Context) {
	service := comment.ListCommentService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//点赞评论
func LikeComment(c *gin.Context) {
	serice := comment.LikeCommentService{}
	res := serice.Like(c)
	c.JSON(200, res)
}

//是否点赞了评论
func IsLikeComment(c *gin.Context) {
	serice := comment.IsLikeCommentService{}
	res := serice.List(c)
	c.JSON(200, res)
}

//不点赞评论
func UnLikeComment(c *gin.Context) {
	serice := comment.UnLikeCommentService{}
	res := serice.UnLike(c)
	c.JSON(200, res)
}

// 删除评论
func DeleteComment(c *gin.Context) {
	service := comment.DeleteCommentService{}
	res := service.Delete(c)
	c.JSON(200, res)
}
