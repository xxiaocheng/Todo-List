package models

import (
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	GroupName string `gorm:"not null"`
	Todos     []Todo
	User      User
	UserId    uint `gorm:"index"`
}

func AutoMigrateGroup() {
	DB.AutoMigrate(&Group{})
}
