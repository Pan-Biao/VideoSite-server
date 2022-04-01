package model

import "gorm.io/gorm"

type Favorites struct {
	gorm.Model
	Name      string
	Collector uint
}
