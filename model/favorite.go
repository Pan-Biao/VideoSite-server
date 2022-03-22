package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	Name      string
	Collector uint
}
