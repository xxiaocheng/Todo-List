package routers

import (
	"github.com/gin-gonic/gin"
	v1 "todoList/routers/api/v1"
)

// No jwt middleware
func AuthRouterRegister(router *gin.RouterGroup) {
	router.POST("/token", v1.UserLogin)
	router.POST("/register", v1.UserRegister)
}

func UserRouterRegister(router *gin.RouterGroup) {
	router.PATCH("/password", v1.ChangeUserPassword)
}

func GroupRouterRegister(router *gin.RouterGroup) {
	router.GET("/", v1.GetGroups)
	router.DELETE("/:group", v1.DeleteOneGroup)
	router.PATCH("/:group", v1.ModifyOneGroupName)
	router.POST("/", v1.CreateOneGroup)
	router.POST("/:group/task", v1.CreateOneTaskWithGroup)
	router.GET("/:group/task", v1.GetTasksWithGroup)
}

func TaskRouterRegister(router *gin.RouterGroup) {
	router.POST("/", v1.CreateOneTaskWithoutGroup)
	router.DELETE("/:task", v1.DeleteOneTask)
	router.GET("/today/", v1.GetTodayTasks)
	router.GET("/default/", v1.GetDefaultGroupTasks)
	router.PATCH("/:task", v1.ModifyTask)
}
