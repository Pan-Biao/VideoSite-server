package comment

import (
	"strconv"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type CreateCommentService struct {
	Comment string `form:"comment" json:"comment" binding:"required,min=1,max=300"`
}

// CreateCommentService 创建收藏夹的服务
func (service *CreateCommentService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)

	vid := c.Param("vid")
	tempid, _ := strconv.Atoi(vid)
	comment := model.Comment{
		Commentator: user.ID,
		Comment:     service.Comment,
		Vid:         uint(tempid),
	}

	if re := funcs.SQLErr(model.DB.Create(&comment).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildComment(comment),
		Msg:  "成功",
	}
}
