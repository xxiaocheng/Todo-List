package models

import (
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	GroupName string `gorm:"not null"`
	Tasks     []Task
	User      User
	UserId    uint `gorm:"index"`
}

// Migrate the `Group` model.
func AutoMigrateGroup() {
	DB.AutoMigrate(&Group{})
}
