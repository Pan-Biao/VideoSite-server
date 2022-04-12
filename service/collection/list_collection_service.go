package collection

import (
	"strconv"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// ListCollectionService 分区列表
type ListCollectionService struct{}

// List 分区列表服务
func (service *ListCollectionService) List(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)

	fid := c.Param("fid")

	collection := []model.Collection{}

	db := model.DB
	id, _ := strconv.Atoi(fid)
	if id != 0 {
		db = db.Where("favorite = ?", id)
	}
	db = db.Where("collector = ?", user.ID)
	if re := funcs.SQLErr(db.Find(&collection).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildCollections(collection),
		Msg:  "成功",
	}
}
