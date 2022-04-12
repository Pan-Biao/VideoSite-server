package comment

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

//不点赞
type IsLikeCommentService struct{}

func (service *IsLikeCommentService) List(c *gin.Context) serializer.Response {
	user := funcs.GetUser(c)

	db := model.DB
	db = db.Where("uid = ?", user.ID)

	commentLikes := []model.CommentLike{}
	if re := funcs.SQLErr(db.Find(&commentLikes).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildLikeComments(commentLikes),
	}
}
