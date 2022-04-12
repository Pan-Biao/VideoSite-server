package collection

import (
	"strconv"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type CreateCollectionService struct {
	FID uint `form:"fid" json:"fid" `
	CID uint `form:"cid" json:"cid" binding:"required"`
}

// CreateCollectionService 收藏的服务
func (service *CreateCollectionService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)

	video, re := funcs.GetVideo(strconv.Itoa(int(service.CID)))
	if re != nil {
		return re.(serializer.Response)
	}

	//判断收藏是否已存在
	collection := model.Collection{
		Collection: video.ID,
		Collector:  user.ID,
	}
	db := model.DB.Where("collector = ? and collection = ?", user.ID, service.CID)
	temp := model.Collection{}
	db.First(&temp)
	if temp.Collector == collection.Collector {
		return serializer.Response{
			Code: 200,
			Data: true,
			Msg:  "收藏已存在",
		}
	}

	//是否进入收藏夹
	if service.FID != 0 {
		favorite := model.Favorites{}
		if re := funcs.SQLErr(model.DB.First(&favorite, service.FID).Error); re != nil {
			return re.(serializer.Response)
		}
		collection.Favorites = service.FID
	}

	if re := funcs.SQLErr(model.DB.Create(&collection).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildCollection(collection),
		Msg:  "成功",
	}
}
