package collection

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// DeleteCollectionService 删除服务
type DeleteCollectionService struct{}

func (service *DeleteCollectionService) Delete(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)

	cid := c.Param("cid")
	collection := model.Collection{}
	count := int64(0)
	db := model.DB.Where("collector = ? and collection = ?", user.ID, cid)
	db.First(&collection).Count(&count)
	if count == 0 {
		return serializer.ReturnData("收藏不存在", true)
	}

	if err := model.DB.Unscoped().Delete(&collection).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("删除成功", true)
}
