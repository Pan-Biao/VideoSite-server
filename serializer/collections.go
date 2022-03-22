package serializer

import "vodeoWeb/model"

// Collections 收藏序列化器
type Collection struct {
	ID        uint  `json:"id"`
	CID       uint  `json:"cid"`
	CreatedAt int64 `json:"created_at"`
}

// BuildCollection 序列化视频
func BuildCollection(item model.Collection) Collection {
	return Collection{
		ID:        item.ID,
		CID:       item.Collection,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildCollections(Collections []model.Collection) []Collection {
	cs := []Collection{}
	for _, item := range Collections {
		cs = append(cs, Collection{
			ID:        item.ID,
			CID:       item.Collection,
			CreatedAt: item.CreatedAt.Unix(),
		})
	}
	return cs
}
