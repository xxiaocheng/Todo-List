package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "todoList/docs"
	"todoList/logs"
	"todoList/routers"
	"todoList/utils/middlewares"
)

// @title TodoList API
// @version 1.0
// @description This is a docs for TodoList.

// @contact.name ChengXiao
// @contact.email cxxlxx0@gmail.com
// @license.name MIT

// @BasePath /api/v1
func main() {
	r := gin.Default()
	r.Use(logs.LoggerToFile())
	r.Use(middlewares.CorsMiddleware())

	// swagger
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	apiV1 := r.Group("/api/v1")

	routers.AuthRouterRegister(apiV1.Group("/auth"))

	//Use JwtAuthMiddleware
	apiV1.Use(middlewares.JwtAuthMiddleware())
	routers.UserRouterRegister(apiV1.Group("/user"))
	routers.GroupRouterRegister(apiV1.Group("/group"))
	routers.TaskRouterRegister(apiV1.Group("/task"))

	_ = r.Run()
}
