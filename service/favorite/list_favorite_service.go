package favorite

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// ListFavoriteService 分区列表
type ListFavoriteService struct{}

// List 分区列表服务
func (service *ListFavoriteService) List(c *gin.Context) serializer.Response {
	favorite := []model.Favorite{}
	err := model.DB.Find(&favorite).Error

	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "视频查询错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildFavorites(favorite),
		Msg:  "成功",
	}
}
