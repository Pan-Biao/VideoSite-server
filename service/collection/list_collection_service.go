package collection

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// ListCollectionService 分区列表
type ListCollectionService struct {
	FID uint `from:"fid" json:"fid" `
}

// List 分区列表服务
func (service *ListCollectionService) List(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	db := model.DB
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}

	collection := []model.Collection{}
	if service.FID != 0 {
		db.Where("favorite = ?", service.FID)
	}
	if err := db.Where("collector = ?", user.ID).Find(&collection).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "收藏查询错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildCollections(collection),
		Msg:  "成功",
	}
}
