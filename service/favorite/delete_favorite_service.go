package favorite

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 删除服务
type DeleteFavoriteService struct{}

// 收藏夹删除
func (service *DeleteFavoriteService) Delete(c *gin.Context) serializer.Response {
	//获取当前用户
	user := funcs.GetUser(c)

	fid := c.Param("fid")
	favorite := model.Favorites{}
	db := model.DB.Where("collector = ? and id = ?", user.ID, fid)
	if re := funcs.SQLErr(db.First(&favorite).Error); re != nil {
		return serializer.Response{
			Code: 200,
			Data: false,
			Msg:  "收藏夹不存在",
		}
	}

	if re := funcs.SQLErr(model.DB.Unscoped().Delete(&favorite).Error); re != nil {
		return re.(serializer.Response)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
		Msg:  "成功",
	}
}
