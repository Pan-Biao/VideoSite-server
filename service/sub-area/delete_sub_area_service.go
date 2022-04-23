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
		return serializer.CheckNoRight()
	}

	said := c.Param("said")
	subArea := model.SubArea{}
	count := int64(0)
	model.DB.First(&subArea, said).Count(&count)
	if count == 0 {
		return serializer.ReturnData("分区不存在", false)
	}

	if err := model.DB.Unscoped().Delete(&subArea).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("删除成功", true)
}
