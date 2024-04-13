package main

import (
	"fmt"

	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	server := gin.Default()
	server.GET("/", utils.GetRoot)
	server.GET("/event/listall", utils.GetAllEvents)
	server.POST("/event/add", utils.SaveEvent)
	server.Run(":8080")

}
