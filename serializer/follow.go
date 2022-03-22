package serializer

import "vodeoWeb/model"

// Follows 关注序列化器
type Follow struct {
	ID        uint  `json:"id"`
	FID       uint  `json:"fid"`
	CreatedAt int64 `json:"created_at"`
}

func BuildFollows(Follows []model.Follow) []Follow {
	fs := []Follow{}
	for _, item := range Follows {
		fs = append(fs, Follow{
			ID:        item.ID,
			FID:       item.Follower,
			CreatedAt: item.CreatedAt.Unix(),
		})
	}
	return fs
}
