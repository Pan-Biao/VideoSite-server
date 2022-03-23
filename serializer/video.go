package serializer

import (
	"vodeoWeb/model"
)

// Video 视频序列化器
type Video struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Info       string `json:"info"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	State      bool   `json:"state"`
	PlayNumber int    `json:"play_number"`
	Uid        uint   `json:"uid"`
	Said       uint   `json:"said"`
	Cover      string `json:"cover"`
	User       User   `json:"user"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	user := model.User{}
	model.DB.First(&user, item.Uid)
	buildUser := BuildUser(user)
	return Video{
		ID:         item.ID,
		Title:      item.Title,
		Info:       item.Info,
		CreatedAt:  item.CreatedAt.Unix(),
		UpdatedAt:  item.UpdatedAt.Unix(),
		State:      item.State,
		PlayNumber: item.PlayNumber,
		Uid:        item.Uid,
		Said:       item.Said,
		Cover:      item.Cover,
		User:       buildUser,
	}
}

// Videos 视频列表序列化器
type Videos struct {
	PageNumber int     `json:"page_number"`
	Number     int     `json:"number"`
	Videos     []Video `json:"videos"`
}

func BuildVideos(items []model.Video, pageNumber int, number int) (videos Videos) {
	vs := []Video{}
	//查询id对应的用户数据
	for _, item := range items {
		video := BuildVideo(item)
		vs = append(vs, video)
	}
	videos.Videos = vs
	videos.PageNumber = pageNumber
	videos.Number = number
	return videos
}
