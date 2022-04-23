package comment

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

//不点赞
type UnLikeCommentService struct{}

func (service *UnLikeCommentService) UnLike(c *gin.Context) serializer.Response {
	cid := c.Param("cid")
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	comment := model.Comment{}
	count := int64(0)
	model.DB.First(&comment, cid).Count(&count)
	if count == 0 {
		return serializer.ReturnData("评论不存在", false)
	}

	commentLike := model.CommentLike{}
	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("cid = ?", cid)
	db.First(&commentLike).Count(&count)
	if count == 0 {
		return serializer.ReturnData("点赞不存在", true)
	}

	if err := model.DB.Delete(&commentLike).Error; err != nil {
		return serializer.DBErr("", err)
	}

	//点赞量-1
	comment.LikeNumber = comment.LikeNumber - 1
	if err := model.DB.Save(&comment).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("取消点赞成功", true)
}
