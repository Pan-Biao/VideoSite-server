package comment

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// DeleteCommentService 分区删除服务
type DeleteCommentService struct{}

// 分区删除
func (service *DeleteCommentService) Delete(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)

	cid := c.Param("cid")
	comment := model.Comment{}
	db := model.DB.Where("commentator = ? and id = ?", user.ID, cid)
	if re := funcs.SQLErr(db.First(&comment).Error); re != nil {
		return serializer.Response{
			Code: 404,
			Data: false,
			Msg:  "评论不存在",
		}
	}

	if re := funcs.SQLErr(model.DB.Unscoped().Delete(&comment).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "成功",
	}
}
