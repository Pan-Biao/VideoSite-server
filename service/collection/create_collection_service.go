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
	CID uint `form:"cid" json:"cid"`
}

// CreateCollectionService 收藏的服务
func (service *CreateCollectionService) Create(c *gin.Context) serializer.Response {
	if service.CID == 0 {
		return serializer.ParamErr("请传入视频id")
	}
	//获取当前用户
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	video := funcs.GetVideo(strconv.Itoa(int(service.CID)))
	if video == (model.Video{}) {
		return serializer.DBErr("", nil)
	}

	//判断收藏是否已存在
	collection := model.Collection{
		Collection: video.ID,
		Collector:  user.ID,
	}
	db := model.DB.Where("collector = ? and collection = ?", user.ID, service.CID)
	temp := model.Collection{}
	count := int64(0)
	db.First(&temp).Count(&count)
	if count != 0 {
		return serializer.ReturnData("收藏已存在", true)
	}

	//是否进入收藏夹
	if service.FID != 0 {
		favorite := model.Favorites{}
		if err := model.DB.First(&favorite, service.FID).Error; err != nil {
			return serializer.DBErr("", err)
		}
		collection.Favorites = service.FID
	}

	if err := model.DB.Create(&collection).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("收藏成功", serializer.BuildCollection(collection))
}
