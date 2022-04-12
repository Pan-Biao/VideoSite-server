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

	comment := model.Comment{}
	if re := funcs.SQLErr(model.DB.First(&comment, cid).Error); re != nil {
		return re.(serializer.Response)
	}

	commentLike := model.CommentLike{}
	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("cid = ?", cid)
	if re := funcs.SQLErr(db.First(&commentLike).Error); re != nil {
		return serializer.Response{
			Code: 200,
			Data: true,
			Msg:  "已取消点赞",
		}
	}

	if re := funcs.SQLErr(model.DB.Delete(&commentLike).Error); re != nil {
		return re.(serializer.Response)
	}

	//点赞量-1
	comment.LikeNumber = comment.LikeNumber - 1
	if re := funcs.SQLErr(model.DB.Save(&comment).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "取消点赞成功",
	}
}
