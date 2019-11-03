package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique_index;not null"`
	PasswordHash string `gorm:"not null"`
	Email        string `gorm:"not null"`
	Todos        []Todo
	Groups       []Group
}

func AutoMigrateUser() {
	DB.AutoMigrate(&User{})
}
