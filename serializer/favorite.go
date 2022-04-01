package serializer

import "vodeoWeb/model"

type Favorite struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// 序列化收藏夹
func BuildFavorites(item model.Favorites) Favorite {
	return Favorite{
		ID:   item.ID,
		Name: item.Name,
	}
}

// 序列化收藏夹列表
func BuildListFavorites(items []model.Favorites) (favorites []Favorite) {
	for _, item := range items {
		favorite := BuildFavorites(item)
		favorites = append(favorites, favorite)
	}
	return favorites
}
