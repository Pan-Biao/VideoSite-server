package favorite

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// DeleteFavoriteService 分区删除服务
type DeleteFavoriteService struct{}

// 分区删除
func (service *DeleteFavoriteService) Delete(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	fid := c.Param("fid")
	favorites := model.Favorites{}
	if err := model.DB.Where("collector = ? and id = ?", user.ID, fid).First(&favorites).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "收藏夹不存在",
			Error: err.Error(),
		}
	}

	if err := model.DB.Delete(&favorites).Error; err != nil {
		return serializer.Response{
			Code:  60003,
			Msg:   "收藏夹删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "成功",
	}
}
