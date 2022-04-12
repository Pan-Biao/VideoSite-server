package favorite

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

type UpdateFavoriteService struct {
	Name string `form:"name" json:"name" binding:"required,min=1,max=12"`
}

// UpdateFavoriteService 收藏夹更新
func (service *UpdateFavoriteService) Update(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)

	fid := c.Param("fid")
	favorites := model.Favorites{}
	db := model.DB.Where("collector = ? and id = ?", user.ID, fid)
	if re := funcs.SQLErr(db.First(&favorites).Error); re != nil {
		return re.(serializer.Response)
	}

	if CheckingFavorite(user, service.Name) {
		return serializer.Response{
			Code: 404,
			Data: false,
			Msg:  "名称重复",
		}
	}
	//更新收藏夹信息
	favorites.Name = service.Name

	if re := funcs.SQLErr(model.DB.Save(&favorites).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildFavorites(favorites),
		Msg:  "成功",
	}
}
