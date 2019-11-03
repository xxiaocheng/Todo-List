package models

import (
	"os"
	"testing"
)

func TestUser(t *testing.T) {
	todoOne := Todo{
		TodoContent: "睡觉",
		HasDeadline: false,
	}
	todoTwo := Todo{
		TodoContent: "吃饭",
		HasDeadline: true,
	}
	group := Group{
		GroupName: "group_one",
		Todos: []Todo{
			todoTwo,
		},
	}
	user := User{
		Username:     "test_name啊",
		PasswordHash: "password",
		Email:        "cxx@gmail.com",
		Todos: []Todo{
			todoOne,
		},
		Groups: []Group{
			group,
		},
	}

	DB.Create(&user)
	DB.Save(&user)
}

func TestMain(m *testing.M) {
	AutoMigrateGroup()
	AutoMigrateTodo()
	AutoMigrateUser()
	exitCode := m.Run()
	// DB.DropTableIfExists(User{})
	// DB.DropTableIfExists(Todo{})
	// DB.DropTableIfExists(Group{})
	os.Exit(exitCode)
}
