package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// err := User.setUserPassword("password0")
// SaveOne(user)
func (user *User) setUserPassword(password string) error {
	if len(password) == 0 {
		errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user.PasswordHash = string(passwordHash)
	return nil
}

// 	if err := User.CheckUserPassword("password0"); err != nil { password error }
func (user *User) CheckUserPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(user.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// 	user, err := FindOneUser(&User{Username: "username0"})
func FindOneUser(condition interface{}, args ...interface{}) (User, error) {
	var userModel User
	err := DB.Where(condition, args...).First(&userModel).Error
	return userModel, err
}

// Find a user by `username`
// user,err:=FindUserByUsername(username)
func FindOneUserByUsername(username string) (user User, err error) {
	user, err = FindOneUser("username = ?", username)
	return
}

func (user *User) SetEmail(email string) error {
	if len(email) == 0 {
		errors.New("email should not be empty!")
	}
	return user.UpdateUserAttr("email", email)
}

//  err := DB.Model(userModel).UpdateUserAttr(User{Username: "wangzitian0"}).Error
func (user *User) UpdateUserAttr(name, value string) error {
	err := DB.Model(user).Update(name, value).Error
	return err
}

// Save update value in database, if the value doesn't have primary key, will insert it
func SaveOne(data interface{}) error {
	err := DB.Save(data).Error
	return err
}

// err:=user.ChangeUsername("new name")
func (user *User) ChangeUsername(newUsername string) error {
	return user.UpdateUserAttr("username", newUsername)
}

// Change password
// err:=user.ChangePassword(newPassword)
func (user *User) ChangePassword(newPassword string) error {
	err := user.setUserPassword(newPassword)
	SaveOne(&user)
	return err
}

// insert a user into database
// user, err := CreateUser(faker.Username(), faker.Password(), faker.Email())
// if err != nil {
// 	panic("user not created")
// }
func CreateUser(username, password, email string) (User, error) {
	user := User{
		Username: username,
		Email:    email,
	}
	user.setUserPassword(password)

	f := DB.NewRecord(user)
	if !f {
		return user, errors.New("User had existed")
	}
	err := DB.Create(&user).Error
	return user, err
}

func FindOneGroupByID(id uint) (Group, error) {
	group := Group{}
	err := DB.First(&group, id).Error
	return group, err
}

// Create a new group
// group,err:=CreateGroup(groupName)
func (user *User) CreateOneGroup(groupName string) (Group, error) {
	group := Group{
		GroupName: groupName,
		User:      *user,
	}
	err := DB.Create(&group).Error
	return group, err
}

// Delete a group if owns by this user.
// err:=user.DeleteGroup(groupID)
func (user *User) DeleteOneGroup(groupID uint) error {
	group := Group{}
	sx := DB.Begin()
	sx.First(&group, groupID)
	if user.ID != group.UserId {
		sx.Commit()
		return errors.New("This group not owns by th user")
	}
	sx.Model(Task{}).Where("group_id = ?", groupID).Update("group_id", 0)
	sx.Delete(&group)
	err := sx.Commit().Error
	return err
}

// err := user.ModifyGroupName(8, "new group name")
// if err != nil {
//		panic("error")
// }
func (user *User) ModifyGroupName(groupID uint, newGroupName string) error {
	group := Group{}
	sx := DB.Begin()
	sx.First(&group, groupID)
	if user.ID != group.UserId {
		sx.Commit()
		return errors.New("This group not owns by th user")
	}
	group.GroupName = newGroupName
	sx.Save(&group)
	err := sx.Commit().Error
	return err
}

// groups, count, err := userModel.PagedGroups(0, 3)
func (user *User) PagedGroups(offset, limit int) ([]Group, int, error) {
	var groups []Group
	sx := DB.Begin()
	sx.Model(&user).Offset(offset).Limit(limit).Related(&groups, "Groups")
	count := sx.Model(&user).Association("Groups").Count()
	err := sx.Commit().Error
	return groups, count, err
}

func (user *User) CreateTaskWithoutGroup(taskContent string) (Task,error) {
	task:=Task{TaskContent:taskContent}
	if task.Deadline == (time.Time{}) {
		task.Deadline = time.Now()
	}
	err := DB.Model(&user).Association("Tasks").Append(&task).Error
	return task,err
}

func (user *User) CreateTaskWithGroup(groupID uint, taskContent string) (Task,error) {
	group:=Group{}
	task:=Task{TaskContent:taskContent}
	if task.Deadline == (time.Time{}) {
		task.Deadline = time.Now()
	}
	sx := DB.Begin()
	sx.First(&group,groupID)
	if group.UserId != user.ID {
		sx.Commit()
		return task,errors.New("This group not owns by the user")
	}
	task.UserId = user.ID
	err := sx.Model(&group).Association("Tasks").Append(&task).Error
	sx.Commit()
	return task,err
}

func (user *User) PagedTodayTasks(offset, limit int) ([]Task, int, error) {
	var tasks []Task
	sx := DB.Begin()
	err := sx.Model(&user).Where("DATEDIFF(deadline,NOW()) = ? ", 0).Limit(limit).Offset(offset).Related(&tasks, "Tasks").Error
	count := sx.Model(&user).Where("DATEDIFF(deadline,NOW()) = ? ", 0).Association("Tasks").Count()
	sx.Commit()
	return tasks, count, err
}

func (user *User) PagedDefaultGroupTasks(offset, limit int) ([]Task, int, error) {
	var tasks []Task
	sx := DB.Begin()
	err := sx.Model(&user).Where("group_id = ?", 0).Limit(limit).Offset(offset).Related(&tasks, "Tasks").Error
	count := sx.Model(&user).Where("group_id = ?", 0).Association("Tasks").Count()
	sx.Commit()
	return tasks, count, err
}

func (user *User) PageTasksWithGroup(groupID uint, offset, limit int) ([]Task, int, error) {
	var group Group
	var tasks []Task
	sx := DB.Begin()
	sx.First(&group, groupID)
	if group.UserId != user.ID {
		sx.Commit()
		return tasks, 0, errors.New("This group not owns by the user")
	}
	err := sx.Model(&group).Limit(limit).Offset(offset).Related(&tasks, "Tasks").Error
	count := sx.Model(&group).Association("Tasks").Count()
	sx.Commit()
	return tasks, count, err
}

func (user *User) DeleteOneTask(taskID uint) error {
	task := Task{}
	DB.First(&task, taskID)
	if task.UserId != user.ID {
		return errors.New("This Task not owns by the user")
	}
	err := DB.Delete(&task).Error
	return err
}

func FindOneTaskById(taskID uint) (Task, error) {
	task := Task{}
	err := DB.First(&task, taskID).Error
	return task, err
}

func (user *User) updateTaskAttr(taskID uint, name string, value interface{}) error {
	task, err := FindOneTaskById(taskID)
	if err != nil {
		return errors.New("Task Not Found")
	}
	if task.UserId != user.ID {
		return errors.New("This Task not owns by the user")
	}
	err = DB.Model(task).Update(name, value).Error
	return err
}

func (user *User) ModifyTaskContent(taskID uint, newContent string) error {
	err := user.updateTaskAttr(taskID, "TaskContent", newContent)
	return err
}

func (user *User) ModifyTaskDeadline(taskID uint, newDeadline time.Time) error {
	err := user.updateTaskAttr(taskID, "Deadline", newDeadline)
	return err
}

func (user *User) ModifyTaskStatus(taskID uint,newStatus bool) error  {
	err := user.updateTaskAttr(taskID, "is_done", newStatus)
	return err
}