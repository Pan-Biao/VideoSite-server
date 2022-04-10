package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	Title string
	Path  string
	Cover string
}
