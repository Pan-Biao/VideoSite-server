package collection

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// ListCollectionService 收藏列表服务
type ListCollectionService struct{}

func (service *ListCollectionService) List(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}
	fid := c.Param("fid")
	collection := []model.Collection{}

	db := model.DB
	db = db.Where("favorites = ?", fid)
	db = db.Where("collector = ?", user.ID)
	if err := db.Find(&collection).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("成功", serializer.BuildCollections(collection))
}
