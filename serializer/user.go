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
	Info         string `json:"info"`
	Token        string `json:"token,omitempty"`
	TokenExpire  int64  `json:"token_expire,omitempty"`
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
		Info:         user.Info,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserToken(user model.User) User {
	return User{
		ID:           user.ID,
		UserName:     user.UserName,
		Nickname:     user.Nickname,
		Status:       user.Status,
		HeadPortrait: user.HeadPortrait,
		CreatedAt:    user.CreatedAt.Unix(),
		UpdatedAt:    user.UpdatedAt.Unix(),
		Root:         user.Root,
		Info:         user.Info,
	}
}

// 用户列表序列化器
type Users struct {
	PageNumber int    `json:"page_number"`
	Number     int    `json:"number"`
	Total      int    `json:"total"`
	Users      []User `json:"users"`
}

// 序列化用户列表
func BuildUsers(items []model.User, pageNumber int, number int, total int) (users Users) {

	for _, item := range items {
		user := BuildUser(item)
		users.Users = append(users.Users, user)
	}
	users.Total = total
	users.Number = number
	users.PageNumber = pageNumber
	return users
}
