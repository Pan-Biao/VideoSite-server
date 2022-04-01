package collection

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

type CreateCollectionService struct {
	FID uint `from:"fid" json:"fid" `
	CID uint `from:"cid" json:"cid" binding:"required"`
}

// CreateCollectionService 收藏的服务
func (service *CreateCollectionService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	video := model.Video{}
	if err := model.DB.First(&video, service.CID).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "无法查找到需要收藏的视频",
			Error: err.Error(),
		}
	}
	collection := model.Collection{}
	if service.FID != 0 {
		favorite := model.Favorites{}
		if err := model.DB.First(&favorite, service.FID).Error; err != nil {
			return serializer.Response{
				Code:  404,
				Msg:   "无法查找到收藏夹",
				Error: err.Error(),
			}
		}
		collection = model.Collection{
			Collection: video.ID,
			Collector:  user.ID,
			Favorites:  service.FID,
		}
	} else {
		collection = model.Collection{
			Collection: video.ID,
			Collector:  user.ID,
		}
	}

	if err := model.DB.Create(&collection).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "收藏视频失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildCollection(collection),
		Msg:  "成功",
	}
}
