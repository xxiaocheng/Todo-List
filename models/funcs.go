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

// 	if err := User.checkUserPassword("password0"); err != nil { password error }
func (user *User) checkUserPassword(password string) error {
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
	sx.Model(Todo{}).Where("group_id = ?", groupID).Update("group_id", 0)
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
	sx.Model(&user).Related(&groups, "Groups").Offset(offset).Limit(limit)
	count := sx.Model(&user).Association("Groups").Count()
	err := sx.Commit().Error
	return groups, count, err
}

func (user *User) CreateTodoWithoutGroup(todo Todo) error {
	if todo.Deadline == (time.Time{}) {
		todo.Deadline = time.Now()
	}
	err := DB.Model(&user).Association("Todos").Append(todo).Error
	return err
}

func (user *User) CreateTodoWithGroup(group Group, todo Todo) error {
	if todo.Deadline == (time.Time{}) {
		todo.Deadline = time.Now()
	}
	sx := DB.Begin()
	sx.Model(&group)
	if group.UserId != user.ID {
		sx.Commit()
		return errors.New("This group not owns by the user")
	}
	todo.UserId = user.ID
	err := sx.Model(&group).Association("Todos").Append(todo).Error
	sx.Commit()
	return err
}

func (user *User) PagedTodayTodos(offset, limit int) ([]Todo, int, error) {
	var todos []Todo
	sx := DB.Begin()
	err := sx.Model(&user).Where("DATEDIFF(deadline,NOW()) = ? ", 0).Limit(limit).Offset(offset).Related(&todos, "Todos").Error
	count := sx.Model(&user).Where("DATEDIFF(deadline,NOW()) = ? ", 0).Association("Todos").Count()
	sx.Commit()
	return todos, count, err
}

func (user *User) PagedDefaultGroupTodos(offset, w int) ([]Todo, int, error) {
	var todos []Todo
	sx := DB.Begin()
	err := sx.Model(&user).Where("group_id = ?", 0).Limit(limit).Offset(offset).Related(&todos, "Todos").Error
	count := sx.Model(&user).Where("group_id = ?", 0).Association("Todos").Count()
	sx.Commit()
	return todos, count, err
}

func (user *User) PageTodosWithGroup(groupID uint, offset, limit int) ([]Todo, int, error) {
	var group Group
	var todos []Todo
	sx := DB.Begin()
	sx.First(&group, groupID)
	if group.UserId != user.ID {
		sx.Commit()
		return todos, 0, errors.New("This group not owns by the user")
	}
	err := sx.Model(&group).Limit(limit).Offset(offset).Related(&todos, "Todos").Error
	count := sx.Model(&group).Association("Todos").Count()
	sx.Commit()
	return todos, count, err
}

func (user *User) DeleteOneTodo(todoID uint) error {
	todo := Todo{}
	DB.First(&todo, todoID)
	if todo.UserId != user.ID {
		return errors.New("This Todo not owns by the user")
	}
	err := DB.Delete(&todo).Error
	return err
}

func FindOneTodoById(todoID uint) (Todo, error) {
	todo := Todo{}
	err := DB.First(&todo, todoID).Error
	return todo, err
}

func (user *User) updateTodoAttr(todoID uint, name string, value interface{}) error {
	todo, err := FindOneTodoById(todoID)
	if err != nil {
		return errors.New("Todo Not Found")
	}
	if todo.UserId != user.ID {
		return errors.New("This Todo not owns by the user")
	}
	err = DB.Model(todo).Update(name, value).Error
	return err
}

func (user *User) ModifyTodoContent(todoID uint, newContent string) error {
	err := user.updateTodoAttr(todoID, "TodoContent", newContent)
	return err
}

func (user *User) ModifyTodoDeadline(todoID uint, newDeadline time.Time) error {
	err := user.updateTodoAttr(todoID, "Deadline", newDeadline)
	return err
}
