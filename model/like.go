package model

import "gorm.io/gorm"

type VideoLike struct {
	gorm.Model
	Uid uint
	Vid uint
}

type CommentLike struct {
	gorm.Model
	Uid uint
	Cid uint
}
