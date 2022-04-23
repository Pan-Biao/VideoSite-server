package favorites

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
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	favorites := []model.Favorites{}
	db := model.DB.Where("collector = ?", user.ID)
	if err := db.Find(&favorites).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("成功", serializer.BuildListFavorites(favorites))
}
