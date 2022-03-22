package collection

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// DeleteCollectionService 分区删除服务
type DeleteCollectionService struct{}

// 分区删除
func (service *DeleteCollectionService) Delete(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	cid := c.Param("cid")
	collection := model.Collection{}
	if err := model.DB.Where("collector = ? and collection = ?", user.ID, cid).First(&collection).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "无法查找到需要删除的收藏",
			Error: err.Error(),
		}
	}

	if err := model.DB.Delete(&collection).Error; err != nil {
		return serializer.Response{
			Code:  60003,
			Msg:   "收藏删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "成功",
	}
}
