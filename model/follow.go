package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	Follower uint
	Fans     uint
}
