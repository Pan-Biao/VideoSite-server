package comment

import (
	"strconv"
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

type CreateCommentService struct {
	Comment string `form:"comment" json:"comment" binding:"required,min=1,max=300"`
}

// CreateCommentService 创建收藏夹的服务
func (service *CreateCommentService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}

	vid := c.Param("vid")
	tempid, _ := strconv.Atoi(vid)
	comment := model.Comment{
		Commentator: user.ID,
		Comment:     service.Comment,
		Vid:         uint(tempid),
	}

	if err := model.DB.Create(&comment).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "评论失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildComment(comment),
		Msg:  "成功",
	}
}
