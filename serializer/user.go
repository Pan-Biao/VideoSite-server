package serializer

import "vodeoWeb/model"

// User 用户序列化器
type User struct {
	ID           uint   `json:"id"`
	UserName     string `json:"user_name"`
	Nickname     string `json:"nickname"`
	Status       string `json:"status"`
	HeadPortrait string `json:"head_portrait"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Root         bool   `json:"root"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:           user.ID,
		UserName:     user.UserName,
		Nickname:     user.Nickname,
		Status:       user.Status,
		HeadPortrait: user.HeadPortrait,
		CreatedAt:    user.CreatedAt.Unix(),
		UpdatedAt:    user.UpdatedAt.Unix(),
		Root:         user.Root,
	}
}

// BuildUsers 序列化用户列表
func BuildUsers(items []model.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
