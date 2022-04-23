package favorites

import (
	"unicode/utf8"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type UpdateFavoriteService struct {
	Name string `form:"name" json:"name"`
}

// UpdateFavoriteService 收藏夹更新
func (service *UpdateFavoriteService) Update(c *gin.Context) serializer.Response {
	if utf8.RuneCountInString(service.Name) < 1 || utf8.RuneCountInString(service.Name) > 12 {
		return serializer.ParamErr("名称长度应为1-12")
	}
	//获取当前用户
	user := funcs.GetUser(c)
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	fid := c.Param("fid")
	favorites := model.Favorites{}
	db := model.DB.Where("collector = ? and id = ?", user.ID, fid)
	count := int64(0)
	db.First(&favorites).Count(&count)
	if count == 0 {
		return serializer.ReturnData("收藏夹不存在", false)
	}

	if CheckingFavorite(user, service.Name) {
		return serializer.ParamErr("名称重复")
	}

	//更新收藏夹信息
	favorites.Name = service.Name
	if err := model.DB.Save(&favorites).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("更新成功", serializer.BuildFavorites(favorites))

}
