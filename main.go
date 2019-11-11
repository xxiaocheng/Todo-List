package main

import (
	"fmt"
	"todoList/logs"
	"todoList/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	token := utils.GenJwtToken("userid")
	fmt.Println(token)

	claims, _ := utils.ParseJwtToken(token)
	fmt.Println((*claims)["exp"])

	r := gin.Default()
	r.Use(logs.LoggerToFile())
	_ = r.Run()
}
