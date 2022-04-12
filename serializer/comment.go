package serializer

import (
	"vodeoWeb/model"
)

type Comment struct {
	ID          uint   `json:"id"`
	Vid         uint   `json:"vid"`
	Comment     string `json:"comment"`
	Commentator uint   `json:"commentator"`
	User        User   `json:"user"`
	CreatedAt   int64  `json:"created_at"`
	LikeNumber  int    `json:"like_number"`
}
type LikeComment struct {
	ID  uint `json:"id"`
	UID uint `json:"uid"`
	CID uint `json:"cid"`
}

// 序列化评论
func BuildComment(item model.Comment) Comment {
	user := model.User{}
	model.DB.Where("id = ?", item.Commentator).First(&user)
	return Comment{
		ID:          item.ID,
		Vid:         item.Vid,
		Comment:     item.Comment,
		Commentator: item.Commentator,
		CreatedAt:   item.CreatedAt.Unix(),
		User:        BuildUser(user),
		LikeNumber:  item.LikeNumber,
	}
}

//点赞信息
func BuildLikeComments(items []model.CommentLike) []LikeComment {
	likeComments := []LikeComment{}
	for _, item := range items {
		likeComment := LikeComment{
			ID:  item.ID,
			UID: item.Uid,
			CID: item.Cid,
		}
		likeComments = append(likeComments, likeComment)
	}
	return likeComments
}

// 序列化评论列表
func BuildComments(items []model.Comment) []Comment {
	comments := []Comment{}
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}
