package main

import (
	"todoList/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(logs.LoggerToFile())
	_ = r.Run()
}
