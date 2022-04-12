package serializer

import "vodeoWeb/model"

// 关注序列化器
type Follow struct {
	ID        uint  `json:"id"`
	CreatedAt int64 `json:"created_at"`
	User      User  `json:"user"`
}

func BuildFollow(item model.Follow, mode string) Follow {
	user := model.User{}
	if mode == "fans" {
		model.DB.First(&user, item.Fans)
	}
	if mode == "follower" {
		model.DB.First(&user, item.Follower)
	}
	follow := Follow{
		ID:        item.ID,
		CreatedAt: item.CreatedAt.Unix(),
		User:      BuildUser(user),
	}
	return follow
}

func BuildFans(Follows []model.Follow) []Follow {
	fs := []Follow{}
	for _, item := range Follows {
		fans := BuildFollow(item, "fans")
		fs = append(fs, fans)
	}
	return fs
}

func BuildFollowers(Follows []model.Follow) []Follow {
	fs := []Follow{}
	for _, item := range Follows {
		follower := BuildFollow(item, "follower")
		fs = append(fs, follower)
	}
	return fs
}
