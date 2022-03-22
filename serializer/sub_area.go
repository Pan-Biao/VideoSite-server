package serializer

import "vodeoWeb/model"

// SubArea 视频序列化器
type SubArea struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// BuildSubArea 序列化视频
func BuildSubArea(item model.SubArea) SubArea {
	return SubArea{
		ID:   item.ID,
		Name: item.Name,
	}
}

// BuildSubAreas 序列化视频列表
func BuildSubAreas(items []model.SubArea) (subAreas []SubArea) {
	for _, item := range items {
		subArea := BuildSubArea(item)
		subAreas = append(subAreas, subArea)
	}
	return subAreas
}
