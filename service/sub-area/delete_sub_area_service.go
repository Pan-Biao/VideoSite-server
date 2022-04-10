package subArea

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// DeleteSubAreaService 分区删除服务
type DeleteSubAreaService struct{}

// 分区删除
func (service *DeleteSubAreaService) Delete(c *gin.Context) serializer.Response {
	if !funcs.CheckRoot(c) {
		return serializer.Response{
			Code: 9999,
			Msg:  "没有权限",
		}
	}

	said := c.Param("said")
	subArea := model.SubArea{}
	if err := model.DB.First(&subArea, said).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "分区不存在",
			Error: err.Error(),
		}
	}

	if err := model.DB.Delete(&subArea).Error; err != nil {
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
