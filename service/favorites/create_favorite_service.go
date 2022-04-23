package favorites

import (
	"unicode/utf8"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type CreateFavoriteService struct {
	Name string `form:"name" json:"name"`
}

// CreateFavoriteService 创建收藏夹的服务
func (service *CreateFavoriteService) Create(c *gin.Context) serializer.Response {
	if utf8.RuneCountInString(service.Name) < 1 || utf8.RuneCountInString(service.Name) > 12 {
		return serializer.ParamErr("名称长度应为1-12")
	}
	//获取当前用户
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	if CheckingFavorite(user, service.Name) {
		return serializer.ParamErr("名称重复")
	}

	favorites := model.Favorites{
		Collector: user.ID,
		Name:      service.Name,
	}

	if err := model.DB.Create(&favorites).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("创建成功", serializer.BuildFavorites(favorites))
}

func CheckingFavorite(user model.User, name string) bool {
	favorites := model.Favorites{}
	model.DB.Where("name = ? and collector = ?", name, user.ID).First(&favorites)
	return favorites.Name != ""
}
