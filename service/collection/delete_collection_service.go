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
	db := model.DB.Where("collector = ? and collection = ?", user.ID, cid)
	if re := funcs.SQLErr(db.First(&collection).Error); re != nil {
		return serializer.Response{
			Code: 200,
			Data: true,
			Msg:  "收藏不存在",
		}
	}

	if re := funcs.SQLErr(model.DB.Unscoped().Delete(&collection).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "成功",
	}
}
