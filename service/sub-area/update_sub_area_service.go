package subArea

import (
	"unicode/utf8"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// UpdateSubAreaService 分区更新服务
type UpdateSubAreaService struct {
	Name string `form:"name" json:"name"`
}

// 分区更新
func (service *UpdateSubAreaService) Update(c *gin.Context) serializer.Response {
	if utf8.RuneCountInString(service.Name) < 2 || utf8.RuneCountInString(service.Name) > 6 {
		return serializer.ParamErr("名称长度应为2-6个字")
	}

	//检测权限
	if !funcs.CheckRoot(c) {
		return serializer.CheckNoRight()
	}

	id := c.Param("id")
	subArea := model.SubArea{}
	count := int64(0)
	model.DB.First(&subArea, id).Count(&count)
	if count == 0 {
		return serializer.ReturnData("分区不存在", false)
	}

	if CheckingSubArea(service.Name) {
		return serializer.ParamErr("名称重复")
	}

	//更新分区信息
	subArea.Name = service.Name

	if err := model.DB.Save(&subArea).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("成功", serializer.BuildSubArea(subArea))
}
