package favorite

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// ListFavoriteService 收藏夹列表
type ListFavoriteService struct{}

// List 收藏夹列表服务
func (service *ListFavoriteService) List(c *gin.Context) serializer.Response {
	favorites := []model.Favorites{}
	uid := c.Param("uid")
	err := model.DB.Where("collector = ?", uid).Find(&favorites).Error

	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "视频查询错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildListFavorites(favorites),
		Msg:  "成功",
	}
}
