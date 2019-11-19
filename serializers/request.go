package serializers

import (
	"strings"
	"time"
	"todoList/models"
)

type AuthorizationHeaderRequest struct {
	Authorization string `header:"Authorization"`
}

func (r *AuthorizationHeaderRequest) StripBearerPrefix() {
	if len(r.Authorization) > 7 && strings.ToUpper(r.Authorization[0:7]) == "BEARER " {
		r.Authorization = r.Authorization[7:]
	}
}

type LoginRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

func (r *LoginRequest) CheckAuth() (bool, error) {
	user, err := models.FindOneUserByUsername(r.Username)
	if err != nil {
		return false, err
	}
	err = user.CheckUserPassword(r.Password)
	if err != nil {
		return false, err
	}
	return true, nil
}

type RegisterUserRequest struct {
	Username string `form:"username" validate:"required,gte=5,lte=16"`
	Password string `form:"password" validate:"required,gte=6,lte=16"`
	Email    string `form:"email" validate:"required,email"`
}

func (r *RegisterUserRequest) Register() (user models.User, err error) {
	user, err = models.CreateUser(r.Username, r.Password, r.Email)
	return
}

type ModifyPasswordRequest struct {
	Password string `form:"password" validate:"required,gte=6,lte=16"`
}

type PageDataRequest struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

func (r *PageDataRequest) ValidDefault() {
	if r.Limit <= 0 {
		r.Limit = 20
	}
	if r.Offset < 0 {
		r.Offset = 0
	}
}

// Modify group name
type ModifyGroupRequest struct {
	NewGroupName string `form:"new_group_name" validate:"required"`
}

// Create a new group
type CreateGroupRequest struct {
	GroupName string `form:"group_name" validate:"required"`
}

// Create a new task
type CreateTaskRequest struct {
	TaskContent string `form:"task_content" validate:"required"`
}

// modify deadline，taskContent,status
type ModifyTaskRequest struct {
	TaskContent string    `form:"task_content"`
	Deadline    time.Time `form:"deadline"`
	IsDone      bool      `form:"is_done"`
}
