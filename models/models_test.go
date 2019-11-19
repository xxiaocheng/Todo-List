package models

import (
	"os"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
)

func TestUser(t *testing.T) {
	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-11-06 22:00:00", time.Local)
	taskOne := Task{
		TaskContent: faker.Name(),
		Deadline:    ti,
	}
	taskTwo := Task{
		TaskContent: faker.Name(),
		Deadline:    time.Now(),
	}

	group := Group{
		GroupName: faker.Name(),
		Tasks: []Task{
			taskTwo,
		},
	}
	user := User{
		Username: faker.Username(),
		Email:    faker.Email(),
		Tasks: []Task{
			taskOne,
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
	user, _ := FindOneUserByUsername("test1")
	for index := 0; index < 50; index++ {
		user.CreateOneGroup(faker.Username())
	}
}

func TestTaskCURD(t *testing.T) {
	user, _ := FindOneUserByUsername("bcAVflI")
	// group, _ := FindOneGroupByID(40)
	// tasks, count, err := user.PageTasksWithGroup(40, 0, 10)
	// if err != nil {
	// 	panic(err)
	// }
	// if count != 33 {
	// 	panic(count)
	// }
	// if len(tasks) != 10 {
	// 	panic(tasks)
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
	AutoMigrateTask()
	AutoMigrateUser()
	exitCode := m.Run()
	// DB.DropTableIfExists(User{})
	// DB.DropTableIfExists(Task{})
	// DB.DropTableIfExists(Group{})
	os.Exit(exitCode)
}
