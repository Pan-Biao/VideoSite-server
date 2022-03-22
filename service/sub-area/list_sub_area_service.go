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
	err := model.DB.Find(&subArea).Error

	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "视频查询错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildSubAreas(subArea),
		Msg:  "成功",
	}
}
