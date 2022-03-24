package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Commentator uint
	Comment     string
	Vid         uint
	LikeNumber  int
}
