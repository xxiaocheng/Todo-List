package models

import (
	"os"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
)

func TestUser(t *testing.T) {
	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-11-06 22:00:00", time.Local)
	todoOne := Todo{
		TodoContent: faker.Name(),
		Deadline:    ti,
	}
	todoTwo := Todo{
		TodoContent: faker.Name(),
		Deadline:    time.Now(),
	}

	group := Group{
		GroupName: faker.Name(),
		Todos: []Todo{
			todoTwo,
		},
	}
	user := User{
		Username: faker.Username(),
		Email:    faker.Email(),
		Todos: []Todo{
			todoOne,
		},
		Groups: []Group{
			group,
		},
	}
	user.setUserPassword(faker.Password())
	DB.Create(&user)
}

func TestUserCURD(t *testing.T) {
	user, err := CreateUser(faker.Username(), faker.Password(), faker.Email())
	if err != nil {
		panic("user not created")
	}
	_, err = CreateUser(user.Username, user.PasswordHash, user.Email)
	// if err != nil {
	// 	panic("user_two not ceated")
	// }
}

func TestGroupCRUD(t *testing.T) {
	user, _ := FindOneUserByUsername("wEDIXuS")
	groups, count, err := user.PagedGroups(0, 3)
	if err != nil {
		panic("sx commit error")
	}
	if len(groups) != 3 {
		panic("lens error")
	}
	if count != 5 {
		panic(count)
	}
}

func TestTodoCURD(t *testing.T) {
	user, _ := FindOneUserByUsername("bcAVflI")
	// group, _ := FindOneGroupByID(40)
	// todos, count, err := user.PageTodosWithGroup(40, 0, 10)
	// if err != nil {
	// 	panic(err)
	// }
	// if count != 33 {
	// 	panic(count)
	// }
	// if len(todos) != 10 {
	// 	panic(todos)
	// }
	err := user.DeleteOneGroup(40)
	if err != nil {
		panic(err)
	}
}

func TestHashIds(t *testing.T) {

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
