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
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	cid := c.Param("cid")
	comment := model.Comment{}
	db := model.DB.Where("commentator = ? and id = ?", user.ID, cid)
	count := int64(0)
	db.First(&comment).Count(&count)
	if count == 0 {
		return serializer.ReturnData("评论不存在", true)
	}

	if err := model.DB.Unscoped().Delete(&comment).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("删除成功", true)
}
