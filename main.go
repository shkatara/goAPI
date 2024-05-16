package main

import (
	"net/http"

	"example.com/api/db"
	"example.com/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.MustRegister(utils.RequestsCounter)
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
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe("127.0.0.1:8000", nil)
	server.Run(":8080")

}
