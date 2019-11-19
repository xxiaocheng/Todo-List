package main

import (
	"todoList/logs"
	"todoList/routers"
	"todoList/utils/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(logs.LoggerToFile())

	apiV1 := r.Group("/api/v1")
	apiV1.POST("/user", routers.UserRegister)
	apiV1.POST("/auth", routers.UserLogin)

	groupApi := apiV1.Group("/group")
	groupApi.Use(middlewares.JwtAuthMiddleware())
	{
		groupApi.GET("/", routers.GetGroups)
		groupApi.DELETE("/:group", routers.DeleteOneGroup)
		groupApi.PATCH("/:group", routers.ModifyOneGroupName)
		groupApi.POST("/", routers.CreateOneGroup)
		groupApi.POST("/:group/todo", routers.CreateOneTaskWithGroup)
		groupApi.GET("/:group/todo", routers.GetTasksWithGroup)
	}

	taskApi := apiV1.Group("/task")
	taskApi.Use(middlewares.JwtAuthMiddleware())
	{
		taskApi.POST("/", routers.CreateOneTaskWithoutGroup)
		taskApi.DELETE("/:task", routers.DeleteOneTask)
		taskApi.GET("/today/", routers.GetTodayTasks)
		taskApi.GET("/default/", routers.GetDefaultGroupTasks)
		taskApi.PATCH("/:task", routers.ModifyTask)
	}

	_ = r.Run()
}
