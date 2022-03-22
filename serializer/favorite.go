package serializer

import "vodeoWeb/model"

type Favorite struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// BuildSubArea 序列化视频
func BuildFavorite(item model.Favorite) Favorite {
	return Favorite{
		ID:   item.ID,
		Name: item.Name,
	}
}

// BuildSubAreas 序列化视频列表
func BuildFavorites(items []model.Favorite) (favorites []Favorite) {
	for _, item := range items {
		favorite := BuildFavorite(item)
		favorites = append(favorites, favorite)
	}
	return favorites
}
