package comment

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// 评论列表
type ListCommentService struct{}

// 评论列表服务
func (service *ListCommentService) List(c *gin.Context) serializer.Response {
	comments := []model.Comment{}
	vid := c.Param("vid")

	if err := model.DB.Where("vid = ?", vid).Find(&comments).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "评论获取失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildComments(comments),
		Msg:  "成功",
	}
}
