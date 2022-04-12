package favorite

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type CreateFavoriteService struct {
	Name string `form:"name" json:"name" binding:"required,min=1,max=12"`
}

// CreateFavoriteService 创建收藏夹的服务
func (service *CreateFavoriteService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)

	if CheckingFavorite(user, service.Name) {
		return serializer.Response{
			Code: 404,
			Data: false,
			Msg:  "名称重复",
		}
	}

	favorites := model.Favorites{
		Collector: user.ID,
		Name:      service.Name,
	}

	if re := funcs.SQLErr(model.DB.Create(&favorites).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildFavorites(favorites),
		Msg:  "成功",
	}
}

func CheckingFavorite(user model.User, name string) bool {
	favorites := model.Favorites{}
	model.DB.Where("name = ? and collector = ?", name, user.ID).First(&favorites)
	return favorites.Name != ""
}
