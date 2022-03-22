package subArea

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// DeleteSubAreaService 分区删除服务
type DeleteSubAreaService struct{}

// 分区删除
func (service *DeleteSubAreaService) Delete(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	if !user.Root {
		return serializer.Response{
			Code: 9999,
			Msg:  "没有权限",
		}
	}
	id := c.Param("id")
	subArea := model.SubArea{}
	err := model.DB.First(&subArea, id).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "分区不存在",
			Error: err.Error(),
		}
	}

	err = model.DB.Delete(&subArea).Error
	if err != nil {
		return serializer.Response{
			Code:  60003,
			Msg:   "分区删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "成功",
	}
}
