package main

import (
	"todoList/logs"
	"todoList/models"

	"github.com/gin-gonic/gin"
	//"todoList/conf"
)

func main() {
	models.AutoMigrateGroup()
	models.AutoMigrateTodo()
	models.AutoMigrateUser()
	r := gin.Default()
	r.Use(logs.LoggerToFile())
	_ = r.Run()
}
