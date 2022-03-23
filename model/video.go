package model

import "gorm.io/gorm"

// Video 视频模型
type Video struct {
	gorm.Model
	Title      string
	Info       string
	Path       string
	State      bool
	PlayNumber int
	Uid        uint
	Said       uint
	Cover      string
}
