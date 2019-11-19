package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	TaskContent string    `gorm:"not null"`
	IsDone      bool      `gorm:"type:boolean;default:false"`
	Deadline    time.Time `gorm:"type:date;default:null"`
	Group       Group
	User        User
	GroupId     uint `gorm:"index"`
	UserId      uint `gorm:"index;not null"`
}

func AutoMigrateTask() {
	DB.AutoMigrate(&Task{})
}
