package comment

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// DeleteCommentService 分区删除服务
type DeleteCommentService struct{}

// 分区删除
func (service *DeleteCommentService) Delete(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	cid := c.Param("cid")
	comment := model.Comment{}
	if err := model.DB.Where("commentator = ? and id = ?", user.ID, cid).First(&comment).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "评论不存在",
			Error: err.Error(),
		}
	}

	if err := model.DB.Delete(&comment).Error; err != nil {
		return serializer.Response{
			Code:  60003,
			Msg:   "评论删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "成功",
	}
}
