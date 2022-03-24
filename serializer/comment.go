package serializer

import "vodeoWeb/model"

type Comment struct {
	ID          uint   `json:"id"`
	Vid         uint   `json:"vid"`
	Comment     string `json:"comment"`
	Commentator uint   `json:"commentator"`
	User        User   `json:"user"`
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
		User:        BuildUser(user),
	}
}

// 序列化评论列表
func BuildComments(items []model.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}
