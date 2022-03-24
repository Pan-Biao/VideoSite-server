package serializer

import "vodeoWeb/model"

type Favorite struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// 序列化收藏夹
func BuildFavorite(item model.Favorite) Favorite {
	return Favorite{
		ID:   item.ID,
		Name: item.Name,
	}
}

// 序列化收藏夹列表
func BuildFavorites(items []model.Favorite) (favorites []Favorite) {
	for _, item := range items {
		favorite := BuildFavorite(item)
		favorites = append(favorites, favorite)
	}
	return favorites
}
