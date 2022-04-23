package subArea

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
)

// ListSubAreaService 分区列表
type ListSubAreaService struct{}

// List 分区列表服务
func (service *ListSubAreaService) List() serializer.Response {
	var subArea []model.SubArea

	if err := model.DB.Find(&subArea).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("成功", serializer.BuildSubAreas(subArea))
}
