package api

import (
	"log"
	"vodeoWeb/service/video"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	service := video.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		log.Println("--------------------11111111111----------------------------------------")
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	service := video.ShowVideoService{}
	res := service.Show(c.Param("vid"))
	c.JSON(200, res)
}

// ListSearchVideo 搜索视频列表接口
func ListSearchVideo(c *gin.Context) {
	service := video.ListSearchVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// UpdateVideo 视频更新接口
func UpdateVideo(c *gin.Context) {
	service := video.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

// DeleteVideo 视频删除接口
func DeleteVideo(c *gin.Context) {
	service := video.DeleteVideoService{}
	res := service.Delete(c)
	c.JSON(200, res)
}

// PlayVideo 视频播放接口
func PlayNumber(c *gin.Context) {
	serice := video.PlayVideoService{}
	serice.Add(c)
}

//点赞
func LikeVideo(c *gin.Context) {
	serice := video.LikeVideoService{}
	serice.Like(c)
}

//不点赞
func UnLikeVideo(c *gin.Context) {
	serice := video.UnLikeVideoService{}
	serice.UnLike(c)
}

//视频封禁
func VideoSuspend(c *gin.Context) {
	service := video.VideoSuspendService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Suspend(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}

//视频解封
func VideoUnseal(c *gin.Context) {
	service := video.VideoUnsealService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Unseal(c)
		c.JSON(200, res)
	} else {
		c.JSON(500, ErrorResponse(err))
	}
}
