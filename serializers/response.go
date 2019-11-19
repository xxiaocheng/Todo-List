package serializers

import (
	"time"
	"todoList/models"
	"todoList/utils/hashID"

	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, CommonResponse{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}

type JwtResponse struct {
	AccessToken string `json:"access_token"`
	ExpireTime  int64  `json:"expire_time"`
}

type PagedDataCommonResponse struct {
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Count  int         `json:"count"`
	Data   interface{} `json:"data"`
}

type GroupResponse struct {
	HashID    string         `json:"id"`
	GroupName string         `json:"group_name"`
	Tasks     []TaskResponse `json:"tasks"`
}

type TaskResponse struct {
	HashID      string    `json:"id"`
	TaskContent string    `json:"task_content"`
	IsDone      bool      `json:"is_done"`
	Deadline    time.Time `json:"deadline"`
}

// Make response for `Group`
func GroupSerializer(group models.Group) GroupResponse {
	var tasks []TaskResponse
	for _,task:=range group.Tasks{
		tasks=append(tasks,TaskSerializer(task))
	}
	return GroupResponse{
		HashID:    hashID.EncodeIDToHash(group.ID),
		GroupName: group.GroupName,
		Tasks:     tasks,
	}
}
// Make response for `Task`
func  TaskSerializer(task models.Task) TaskResponse {
	return TaskResponse{
		HashID:      hashID.EncodeIDToHash(task.ID),
		TaskContent: task.TaskContent,
		IsDone:      task.IsDone,
		Deadline:    task.Deadline,
	}
}

func SerializerGroupsFromModel(groupModels []models.Group) []GroupResponse{
	var groupResponses []GroupResponse
	for _,groupModel :=range groupModels {
		groupResponses=append(groupResponses,GroupSerializer(groupModel))
	}
	return groupResponses
}

func SerializerTasksFromModel(taskModels []models.Task) []TaskResponse{
	var taskResponses []TaskResponse
	for _,taskModel :=range taskModels{
		taskResponses=append(taskResponses,TaskSerializer(taskModel))
	}
	return taskResponses
}
