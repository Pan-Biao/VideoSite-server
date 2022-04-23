package comment

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 是否点赞
type IsLikeCommentService struct{}

func (service *IsLikeCommentService) List(c *gin.Context) serializer.Response {
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	db := model.DB
	db = db.Where("uid = ?", user.ID)
	commentLikes := []model.CommentLike{}

	if err := db.Find(&commentLikes).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("", serializer.BuildLikeComments(commentLikes))
}
