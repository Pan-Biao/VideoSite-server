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
	count := int64(0)
	db.First(&collection).Count(&count)
	if count == 0 {
		return serializer.ReturnData("未收藏", false)
	}

	return serializer.ReturnData("已收藏", true)

}
