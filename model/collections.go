package model

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Collector  uint
	Collection uint
	Favorites  uint
}
