package favorites

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
	if user == (model.User{}) {
		return serializer.DBErr("", nil)
	}

	fid := c.Param("fid")
	favorite := model.Favorites{}
	db := model.DB.Where("collector = ? and id = ?", user.ID, fid)
	count := int64(0)
	db.First(&favorite).Count(&count)
	if count == 0 {
		return serializer.ReturnData("收藏夹不存在", false)
	}

	if err := model.DB.Unscoped().Delete(&favorite).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("删除成功", true)
}
