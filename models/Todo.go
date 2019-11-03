package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	TodoContent string `gorm:"not null"`
	IsDone      bool   `gorm:"type:boolean;default:false"`
	HasDeadline bool   `gorm:"type:boolean"`
	Group       Group
	User        User
	GroupId     uint `gorm:"index"`
	UserId      uint `gorm:"index;not null"`
}

func AutoMigrateTodo() {
	DB.AutoMigrate(&Todo{})
}
