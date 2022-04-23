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
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}
	//查询评论是否存在
	comment := model.Comment{}
	count := int64(0)
	model.DB.First(&comment, cid).Count(&count)
	if count == 0 {
		return serializer.ReturnData("评论不存在", false)
	}

	//创建CommentLike结构
	commentLike := model.CommentLike{
		Uid: user.ID,
		Cid: comment.ID,
	}

	//查询点赞是否已存在
	temp := model.CommentLike{}
	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("cid = ?", comment.ID)
	db.First(&temp).Count(&count)
	if count != 0 {
		return serializer.ReturnData("已经点赞了", true)
	}

	//存进数据库
	if err := model.DB.Create(&commentLike).Error; err != nil {
		return serializer.DBErr("", err)
	}

	//播放量加1
	comment.LikeNumber = comment.LikeNumber + 1
	if err := model.DB.Save(&comment).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("点赞成功", true)
}
