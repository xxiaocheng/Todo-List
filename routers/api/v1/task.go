package v1

import (
	"net/http"
	"time"
	"todoList/models"
	"todoList/serializers"
	"todoList/utils/hashID"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Create one task with group
// @Tags Group
// @Accept multipart/form-data
// @Produce  json
// @Param group path string true "group hashID"
// @param task_content formData string true "task content"
// @Param Authorization header string true "Bearer"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /group/{group}/task [post]
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

// @Summary Create one task without group
// @Tags Task
// @Accept multipart/form-data
// @Produce  json
// @param task_content formData string true "task content"
// @Param Authorization header string true "Bearer"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /task [post]
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

// @Summary Get Today`s tasks
// @Tags Task
// @Accept multipart/form-data
// @Produce  json
// @param offset query integer false "offset"
// @param limit query integer false "limit"
// @Param Authorization header string true "Bearer"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /task/today/ [get]
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

// @Summary Get default`s tasks
// @Tags Task
// @Accept multipart/form-data
// @Produce  json
// @param offset query integer false "offset"
// @param limit query integer false "limit"
// @Param Authorization header string true "Bearer"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /task/default/ [get]
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

// @Summary Get some group`s tasks
// @Tags Group
// @Accept multipart/form-data
// @Produce  json
// @param offset query integer false "offset"
// @param limit query integer false "limit"
// @Param Authorization header string true "Bearer"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /group/{group}/task [get]
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

// @Summary Delete one task
// @Tags Task
// @Accept multipart/form-data
// @Produce  json
// @param task path string true "task hashID"
// @Param Authorization header string true "Bearer"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /task/{task} [delete]
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

// @Summary Modify one task
// @Tags Task
// @Accept multipart/form-data
// @Produce  json
// @Param Authorization header string true "Bearer"
// @param task_content formData string false "task content"
// @param deadline formData string false "task deadline"
// @param is_done formData boolean false "task status"
// @param task path string true "task hashID"
// @Success 200 {object} serializers.CommonResponse "OK"
// @Failure 400 {object} serializers.CommonResponse "FAIL"
// @Router /task [get]
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
