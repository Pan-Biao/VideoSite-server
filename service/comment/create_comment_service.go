package comment

import (
	"strconv"
	"unicode/utf8"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type CreateCommentService struct {
	Comment string `form:"comment" json:"comment"`
}

// CreateCommentService 创建收藏夹的服务
func (service *CreateCommentService) Create(c *gin.Context) serializer.Response {
	if utf8.RuneCountInString(service.Comment) < 1 && utf8.RuneCountInString(service.Comment) > 100 {
		return serializer.ParamErr("评论长度应为1-100")
	}
	//获取当前用户
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	vid := c.Param("vid")
	tempid, _ := strconv.Atoi(vid)
	comment := model.Comment{
		Commentator: user.ID,
		Comment:     service.Comment,
		Vid:         uint(tempid),
	}

	if err := model.DB.Create(&comment).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("评论成功", serializer.BuildComment(comment))
}
