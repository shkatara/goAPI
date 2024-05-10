package main

import (
	"example.com/api/db"
	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.CreateEventTable()
	server := gin.Default()
	server.GET("/", utils.GetRoot)
	server.GET("/event/all/", utils.GetAllEvents)
	server.POST("/event/add/", utils.AddEvent)
	server.GET("/event/fetch/:id", utils.FetchEvent)
	server.POST("/event/delete/:id", utils.DeleteEvent)
	server.POST("/event/update/:id", utils.UpdateEvent)
	server.GET("/redirect", utils.Redirect)
	server.NoRoute(utils.NoRouteHandler)
	server.Run(":8080")

}
