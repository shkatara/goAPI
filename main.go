package main

import (
	"net/http"

	"example.com/api/controller/events"
	"example.com/api/controller/users"
	"example.com/api/db"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func initialize_api() {
	server := gin.Default()
	server.GET("/", events.GetRoot)
	server.GET("/api/event/all/", events.GetAllEvents)
	server.POST("/api/event/add/", events.AddEvent)
	server.GET("/api/event/fetch/:id", events.FetchEvent)
	server.POST("/api/event/delete/:id", events.DeleteEvent)
	server.POST("/api/event/update/:id", events.UpdateEvent)
	server.GET("/api/redirect", events.Redirect)
	server.POST("/api/user/signup/", users.Signup)
	server.POST("/api/user/login/", users.Login)
	server.NoRoute(events.NoRouteHandler)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe("127.0.0.1:8000", nil)
	server.Run(":8080")
}

func main() {
	prometheus.MustRegister(events.RequestsCounter)
	db.InitDB()
	db.CreateEventTable()
	initialize_api()

}
