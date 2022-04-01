package favorite

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

type CreateFavoriteService struct {
	Name string `form:"name" json:"name" binding:"required,min=1,max=12"`
}

// CreateFavoriteService 创建收藏夹的服务
func (service *CreateFavoriteService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	if CheckingFavorite(user, service.Name) {
		return serializer.Response{
			Code: 404,
			Msg:  "名称重复",
		}
	}

	favorites := model.Favorites{
		Collector: user.ID,
		Name:      service.Name,
	}

	if err := model.DB.Create(&favorites).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "收藏夹创建失败",
			Error: err.Error(),
		}
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
