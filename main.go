package main

import (
	"github.com/gin-gonic/gin"
	"todoList/logs"
	"todoList/routers"
	"todoList/utils/middlewares"
)

func main() {
	r := gin.Default()
	r.Use(logs.LoggerToFile())

	apiV1 := r.Group("/api/v1")

	routers.AuthRouterRegister(apiV1.Group("/auth"))

	//Use JwtAuthMiddleware
	apiV1.Use(middlewares.JwtAuthMiddleware())
	routers.UserRouterRegister(apiV1.Group("/user"))
	routers.GroupRouterRegister(apiV1.Group("/group"))
	routers.TaskRouterRegister(apiV1.Group("/task"))

	_ = r.Run()
}
