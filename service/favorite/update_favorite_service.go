package favorite

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

type UpdateFavoriteService struct {
	Name string `form:"name" json:"name" binding:"required,min=1,max=12"`
}

// UpdateFavoriteService 收藏夹更新
func (service *UpdateFavoriteService) Update(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	fid := c.Param("fid")
	favorite := model.Favorite{}
	if err := model.DB.Where("collector = ? and id = ?", user.ID, fid).First(&favorite).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "收藏夹不存在",
			Error: err.Error(),
		}
	}
	if CheckingFavorite(user, service.Name) {
		return serializer.Response{
			Code: 404,
			Msg:  "名称重复",
		}
	}
	//更新收藏夹信息
	favorite.Name = service.Name

	if err := model.DB.Save(&favorite).Error; err != nil {
		return serializer.Response{
			Code:  60002,
			Msg:   "收藏夹更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildFavorite(favorite),
		Msg:  "成功",
	}
}
