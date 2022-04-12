package serializer

import (
	"vodeoWeb/model"
)

// Collections 收藏序列化器
type Collection struct {
	ID        uint  `json:"id"`
	CID       uint  `json:"cid"`
	CreatedAt int64 `json:"created_at"`
	Video     Video `json:"video"`
}

// BuildCollection 序列化收藏
func BuildCollection(item model.Collection) Collection {
	video := model.Video{}
	model.DB.First(&video, item.Collection)

	return Collection{
		ID:        item.ID,
		CID:       item.Collection,
		CreatedAt: item.CreatedAt.Unix(),
		Video:     BuildVideoOne(video),
	}
}

func BuildCollections(Collections []model.Collection) []Collection {
	cs := []Collection{}
	for _, item := range Collections {
		collection := BuildCollection(item)
		cs = append(cs, collection)
	}
	return cs
}
