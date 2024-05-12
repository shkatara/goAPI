package utils

import (
	"database/sql"
	"net/http"

	"example.com/api/db"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type Event struct {
	EventName  string `json:"event_name" binding:"required"`
	EventOwner string `json:"event_owner" binding:"required"`
	EventID    int    `json:"event_id"`
}

var listOfEvents = []map[string]string{
	{
		"event_name":  "A Go-lang life",
		"event_owner": "Shubham",
	},
}

var (
	GetAllEventsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "getallevent_requests_total",
			Help: "Counter to expose GET method to GetAllEvents handler.",
		},
		[]string{"method"},
	)
	AddEventCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "addevent_requests_total",
			Help: "Counter to expose POST method to AddEvent API.",
		},
		[]string{"method"},
	)
	FetchEventCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fetchevent_requests_total",
			Help: "Counter to expose POST method to FetchEvent API.",
		},
		[]string{"method"},
	)
	DeleteEventCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "deleteevent_requests_total",
			Help: "Counter to expose POST method to DeleteEvent API.",
		},
		[]string{"method"},
	)
	UpdateEventCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "updateevent_requests_total",
			Help: "Counter to expose POST method to UpdateEvent API.",
		},
		[]string{"method"},
	)
	RedirectCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "redirect_requests_total",
			Help: "Counter to expose GET method to Redirect API.",
		},
		[]string{"method"},
	)
	NoRouteHandlerCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "noroutehandler_requests_total",
			Help: "Counter to expose calls to NoRouteHandler API",
		},
		[]string{"method"},
	)
)

var events = []Event{
	{EventName: "A Go-lang life", EventOwner: "Shubham", EventID: 1},
	{EventName: "There goes another", EventOwner: "Shubham", EventID: 2},
}

func Redirect(c *gin.Context) {
	RedirectCounter.WithLabelValues("GET").Inc()
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}

func NoRouteHandler(c *gin.Context) {
	NoRouteHandlerCounter.WithLabelValues("GET").Inc()
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Route not found",
	})
}

func GetAllEvents(c *gin.Context) {
	GetAllEventsCounter.WithLabelValues("GET").Inc()
	var events_data []Event
	var event Event
	result, err := db.DB.Query("SELECT id,event_title,event_owner FROM events")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No events found",
		})
	}
	defer result.Close()
	for result.Next() {
		result.Scan(&event.EventID, &event.EventName, &event.EventOwner)
		events_data = append(events_data, event)
	}

	c.JSON(http.StatusOK, gin.H{
		"events_list": events_data,
	})
}

func AddEvent(c *gin.Context) {
	AddEventCounter.WithLabelValues("POST").Inc()
	//event_id := rand.Intn(100000)
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	sql_statement := "INSERT INTO events (event_title, event_owner) VALUES (?, ?)"
	CheckError(err)
	_, err = db.DB.Exec(sql_statement, jsonData.EventName, jsonData.EventOwner)
	CheckError(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "Event added",
	})
}

func FetchEvent(c *gin.Context) {
	FetchEventCounter.WithLabelValues("GET").Inc()
	post_id := c.Param("id")
	var jsonData Event
	var event Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	row := db.DB.QueryRow("SELECT id,event_title,event_owner FROM events where id = ?", post_id)
	err_scan := row.Scan(&event.EventID, &event.EventName, &event.EventOwner)
	if sql.ErrNoRows != err_scan {
		c.JSON(http.StatusOK, gin.H{
			"Event Name":  event.EventName,
			"Event Owner": event.EventOwner,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
	}
}

func DeleteEvent(c *gin.Context) {
	DeleteEventCounter.WithLabelValues("POST").Inc()
	post_id := c.Param("id")
	var jsonData Event
	var event Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	row := db.DB.QueryRow("SELECT id,event_title,event_owner FROM events where id = ?", post_id)
	err_scan := row.Scan(&event.EventID, &event.EventName, &event.EventOwner)
	if sql.ErrNoRows == err_scan {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
	} else {
		sql_statement := "DELETE FROM events where id = ?"
		db.DB.Exec(sql_statement, post_id)
		c.JSON(http.StatusOK, gin.H{
			"message": "Event Deleted",
		})
	}

}

func UpdateEvent(c *gin.Context) {
	UpdateEventCounter.WithLabelValues("POST").Inc()
	post_id := c.Param("id")
	var jsonData Event
	var event Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	row := db.DB.QueryRow("SELECT id,event_title,event_owner FROM events where id = ?", post_id)
	err_scan := row.Scan(&event.EventID, &event.EventName, &event.EventOwner)
	if sql.ErrNoRows == err_scan {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
	} else {
		sql_statement := "Update events SET event_title = ?, event_owner = ? where id = ?"
		db.DB.Exec(sql_statement, jsonData.EventName, jsonData.EventOwner, post_id)
		c.JSON(http.StatusOK, gin.H{
			"message": "Event Updated",
		})
	}
}

func GetRoot(c *gin.Context) {
	returnData := map[string]string{"name": "Shubham"}
	c.JSON(http.StatusOK, gin.H{
		"name":   returnData["name"],
		"events": listOfEvents[0]["event_name"],
	})
}
