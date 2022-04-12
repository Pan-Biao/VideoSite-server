package collection

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type IsCollectionService struct{}

func (service *IsCollectionService) Is(c *gin.Context) serializer.Response {
	user := funcs.GetUser(c)
	collection := model.Collection{}
	cid := c.Param("cid")

	db := model.DB
	db = db.Where("collector = ?", user.ID)
	db = db.Where("Collection = ?", cid)
	if re := db.First(&collection).Error; re != nil {
		return serializer.Response{
			Code: 200,
			Data: false,
			Msg:  "未收藏",
		}
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "已收藏",
	}

}
