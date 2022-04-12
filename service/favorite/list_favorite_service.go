package favorite

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// ListFavoriteService 收藏夹列表
type ListFavoriteService struct{}

// List 收藏夹列表服务
func (service *ListFavoriteService) List(c *gin.Context) serializer.Response {
	user := funcs.GetUser(c)

	favorites := []model.Favorites{}
	db := model.DB.Where("collector = ?", user.ID)
	if re := funcs.SQLErr(db.Find(&favorites).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildListFavorites(favorites),
		Msg:  "成功",
	}
}
