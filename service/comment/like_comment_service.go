package comment

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 视频点赞服务
type LikeCommentService struct{}

func (service *LikeCommentService) Like(c *gin.Context) serializer.Response {
	cid := c.Param("cid")
	user := funcs.GetUser(c)
	//查询评论是否存在
	comment := model.Comment{}
	if re := funcs.SQLErr(model.DB.First(&comment, cid).Error); re != nil {
		return re.(serializer.Response)
	}
	//创建CommentLike结构
	commentLike := model.CommentLike{
		Uid: user.ID,
		Cid: comment.ID,
	}
	//查询是否已存在
	temp := model.CommentLike{}
	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("cid = ?", comment.ID)
	db.First(&temp)
	if commentLike.Uid == temp.Uid {
		return serializer.Response{
			Code: 200,
			Data: true,
			Msg:  "已经点赞过了",
		}
	}
	//存进数据库
	if re := funcs.SQLErr(model.DB.Create(&commentLike).Error); re != nil {
		return re.(serializer.Response)
	}

	//播放量加1
	comment.LikeNumber = comment.LikeNumber + 1
	if re := funcs.SQLErr(model.DB.Save(&comment).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "点赞成功",
	}
}
