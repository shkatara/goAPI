package main

import (
	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", utils.GetRoot)
	server.GET("/event/all/", utils.GetAllEvents)
	server.POST("/event/add/", utils.AddEvent)
	server.GET("/event/fetch/", utils.FetchEvent)
	server.POST("/event/delete/", utils.DeleteEvent)
	server.Run(":8080")

}
