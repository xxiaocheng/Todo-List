package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Username     string `gorm:"unique_index;unique;not null"`
	PasswordHash string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Tasks        []Task
	Groups       []Group
}

// Migrate the `User` model
func AutoMigrateUser() {
	DB.AutoMigrate(&User{})
}
