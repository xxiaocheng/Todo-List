package routers

import (
	"net/http"
	"time"
	"todoList/models"
	"todoList/serializers"
	"todoList/utils/hashID"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// url="/:group/task/"
func CreateOneTaskWithGroup(c *gin.Context) {
	groupHashID := c.Param("group")
	groupID := hashID.DecodeHashToID(groupHashID)
	userModel := (c.MustGet("userModel")).(models.User)
	r := serializers.CreateTaskRequest{}
	appG := serializers.Gin{C: c}
	if err := c.ShouldBind(&r); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	validate := validator.New()
	err := validate.Struct(&r)
	if err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	task, err := userModel.CreateTaskWithGroup(groupID, r.TaskContent)
	if err != nil {
		appG.Response(http.StatusForbidden, serializers.ErrorPutTaskFail, nil)
		return
	}
	appG.Response(http.StatusCreated, serializers.Success, serializers.TaskSerializer(task))
	return
}

func CreateOneTaskWithoutGroup(c *gin.Context) {
	userModel := (c.MustGet("userModel")).(models.User)
	r := serializers.CreateTaskRequest{}
	appG := serializers.Gin{C: c}
	if err := c.ShouldBind(&r); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	validate := validator.New()
	err := validate.Struct(&r)
	if err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	task, err := userModel.CreateTaskWithoutGroup(r.TaskContent)
	if err != nil {
		appG.Response(http.StatusForbidden, serializers.ErrorPutTaskFail, nil)
		return
	}
	appG.Response(http.StatusCreated, serializers.Success, serializers.TaskSerializer(task))
	return

}

func GetTodayTasks(c *gin.Context) {
	pageDataRequest := serializers.PageDataRequest{}
	appG := serializers.Gin{C: c}
	if err := c.ShouldBindQuery(&pageDataRequest); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	pageDataRequest.ValidDefault()
	userModel := (c.MustGet("userModel")).(models.User)
	todayTasks, count, err := userModel.PagedTodayTasks(pageDataRequest.Offset, pageDataRequest.Limit)
	if err != nil {
		appG.Response(http.StatusInternalServerError, serializers.ErrorGetGroupFail, nil)
		return
	}
	cr := serializers.PagedDataCommonResponse{Offset: pageDataRequest.Offset,
		Limit: pageDataRequest.Limit,
		Count: count,
		Data:  serializers.SerializerTasksFromModel(todayTasks),
	}
	appG.Response(http.StatusOK, serializers.Success, cr)
}

func GetDefaultGroupTasks(c *gin.Context) {
	pageDataRequest := serializers.PageDataRequest{}
	appG := serializers.Gin{C: c}
	if err := c.ShouldBindQuery(&pageDataRequest); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	pageDataRequest.ValidDefault()
	userModel := (c.MustGet("userModel")).(models.User)
	todayTasks, count, err := userModel.PagedDefaultGroupTasks(pageDataRequest.Offset, pageDataRequest.Limit)
	if err != nil {
		appG.Response(http.StatusInternalServerError, serializers.ErrorGetGroupFail, nil)
		return
	}
	cr := serializers.PagedDataCommonResponse{Offset: pageDataRequest.Offset,
		Limit: pageDataRequest.Limit,
		Count: count,
		Data:  serializers.SerializerTasksFromModel(todayTasks),
	}
	appG.Response(http.StatusOK, serializers.Success, cr)
}

func GetTasksWithGroup(c *gin.Context) {
	pageDataRequest := serializers.PageDataRequest{}
	appG := serializers.Gin{C: c}
	if err := c.ShouldBindQuery(&pageDataRequest); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	pageDataRequest.ValidDefault()
	groupHashID := c.Param("group")
	groupID := hashID.DecodeHashToID(groupHashID)
	userModel := (c.MustGet("userModel")).(models.User)
	tasks, count, err := userModel.PageTasksWithGroup(groupID, pageDataRequest.Offset, pageDataRequest.Limit)
	if err != nil {
		appG.Response(http.StatusInternalServerError, serializers.ErrorGetGroupFail, nil)
		return
	}
	cr := serializers.PagedDataCommonResponse{Offset: pageDataRequest.Offset,
		Limit: pageDataRequest.Limit,
		Count: count,
		Data:  serializers.SerializerTasksFromModel(tasks),
	}
	appG.Response(http.StatusOK, serializers.Success, cr)
}

func DeleteOneTask(c *gin.Context) {
	taskHashID := c.Param("task")
	taskID := hashID.DecodeHashToID(taskHashID)
	userModel := (c.MustGet("userModel")).(models.User)
	err := userModel.DeleteOneTask(taskID)
	appG := serializers.Gin{C: c}
	if err != nil {
		appG.Response(http.StatusForbidden, serializers.ErrorDeleteTaskFail, nil)
	} else {
		appG.Response(http.StatusOK, serializers.Success, nil)
	}
}

func ModifyTask(c *gin.Context) {
	appG := serializers.Gin{C: c}
	r := serializers.ModifyTaskRequest{
		TaskContent: "",
		Deadline:    time.Time{},
		IsDone:      false,
	}
	userModel := (c.MustGet("userModel")).(models.User)
	if err := c.ShouldBind(&r); err != nil {
		appG.Response(http.StatusBadRequest, serializers.InvalidParams, nil)
		return
	}
	taskID := hashID.DecodeHashToID(c.Param("task"))
	if r.TaskContent != "" {
		err := userModel.ModifyTaskContent(taskID, r.TaskContent)
		if err != nil {
			appG.Response(http.StatusForbidden, serializers.ErrorEditTaskFail, nil)
			return
		}
	} else if r.Deadline != (time.Time{}) {
		err := userModel.ModifyTaskDeadline(taskID, r.Deadline)
		if err != nil {
			appG.Response(http.StatusForbidden, serializers.ErrorEditTaskFail, nil)
			return
		}
	} else {
		err := userModel.ModifyTaskStatus(taskID, r.IsDone)
		if err != nil {
			appG.Response(http.StatusForbidden, serializers.ErrorEditTaskFail, nil)
			return
		}
	}
	appG.Response(http.StatusOK, serializers.Success, nil)
}
