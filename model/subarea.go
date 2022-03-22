package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type SubArea struct {
	gorm.Model
	Name string
}
