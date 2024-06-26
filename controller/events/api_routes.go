package events

import (
	"database/sql"
	"fmt"
	"net/http"

	"example.com/api/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prometheus/client_golang/prometheus"
)

type Event struct {
	EventTitle   string `json:"event_title" binding:"required"`
	EventContent string `json:"event_content" binding:"required"`
	EventOwnerID string `json:"event_owner_name"`
	EventID      int    `json:"event_id"`
}

var listOfEvents = []map[string]string{
	{
		"event_name":  "A Go-lang life",
		"event_owner": "Shubham",
		"event_title": "I have started liking golang",
	},
}
var hmacSampleSecret = []byte("dummytestSecret")

var (
	RequestsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Counter to expose GET method to GetAllEvents handler.",
		},
		[]string{"method", "endpoint"},
	)
)

var events = []Event{
	{EventTitle: "A Go-lang life", EventOwnerID: "Shubham", EventID: 1},
	{EventTitle: "There goes another", EventOwnerID: "Shubham", EventID: 2},
}

func Redirect(c *gin.Context) {
	RequestsCounter.WithLabelValues("GET", "redirect").Inc()
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}

func NoRouteHandler(c *gin.Context) {
	RequestsCounter.WithLabelValues("GET", "noroute").Inc()
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Route not found",
	})
}

func GetAllEvents(c *gin.Context) {
	RequestsCounter.WithLabelValues("GET", "all").Inc()
	var events_data []Event
	var event Event
	AuthorizationToken := c.GetHeader("Authorization")
	token, isvalid := IsValid(AuthorizationToken)
	if isvalid {
		claims, _ := token.Claims.(jwt.MapClaims)
		sql_statement := fmt.Sprintf("SELECT events.event_title, events.event_content FROM events where event_owner_name = '%s'", claims["username"])
		result, _ := db.DB.Query(sql_statement)
		defer result.Close()
		for result.Next() {
			result.Scan(&event.EventTitle, &event.EventContent)
			events_data = append(events_data, event)
		}
		if len(events_data) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"events_list": "No Events Found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"events_list": events_data,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization Failed",
		})
	}
}

func AddEvent(c *gin.Context) {
	RequestsCounter.WithLabelValues("POST", "add").Inc()
	var jsonData Event
	c.ShouldBindJSON(&jsonData)
	AuthorizationToken := c.GetHeader("Authorization")
	token, isvalid := IsValid(AuthorizationToken)
	if isvalid {
		claims, _ := token.Claims.(jwt.MapClaims)
		sql_statement := "INSERT INTO events (event_title, event_content, event_owner_name ) VALUES (?, ?, ?)"
		_, err := db.DB.Exec(sql_statement, jsonData.EventTitle, jsonData.EventContent, claims["username"])
		CheckError(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "Event added",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization Failed",
		})
	}
}

func FetchEvent(c *gin.Context) {
	RequestsCounter.WithLabelValues("GET", "fetch").Inc()
	post_id := c.Param("id")
	var jsonData Event
	var event Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	AuthorizationToken := c.GetHeader("Authorization")
	token, isvalid := IsValid(AuthorizationToken)
	if isvalid {
		claims, _ := token.Claims.(jwt.MapClaims)
		sql_statement := fmt.Sprintf("SELECT events.event_title, events.event_content FROM events where event_owner_name = '%s' and event_id = %s", claims["username"], post_id)
		row := db.DB.QueryRow(sql_statement)
		err_scan := row.Scan(&event.EventTitle, &event.EventContent)
		if sql.ErrNoRows != err_scan {
			c.JSON(http.StatusOK, gin.H{
				"Event Title":   event.EventTitle,
				"Event Content": event.EventContent,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Event not found",
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization Failed",
		})
	}
}

func DeleteEvent(c *gin.Context) {
	RequestsCounter.WithLabelValues("POST", "delete").Inc()
	post_id := c.Param("id")
	var jsonData Event
	var event Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	AuthorizationToken := c.GetHeader("Authorization")
	token, isvalid := IsValid(AuthorizationToken)
	if isvalid {
		claims, _ := token.Claims.(jwt.MapClaims)
		delete_sql_statement := fmt.Sprintf("DELETE FROM events where event_owner_name = '%s' and event_id = %s", claims["username"], post_id)
		fetch_sql_statement := fmt.Sprintf("SELECT events.event_title, events.event_content FROM events where event_owner_name = '%s' and event_id = %s", claims["username"], post_id)
		row := db.DB.QueryRow(fetch_sql_statement)
		err_scan := row.Scan(&event.EventTitle, &event.EventContent)
		if sql.ErrNoRows == err_scan {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Event not found",
			})
		} else {
			db.DB.Exec(delete_sql_statement)
			c.JSON(http.StatusOK, gin.H{
				"message": "Event Deleted",
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization Failed",
		})
	}
}

func UpdateEvent(c *gin.Context) {
	RequestsCounter.WithLabelValues("POST", "update").Inc()
	post_id := c.Param("id")
	var jsonData Event
	var event Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	AuthorizationToken := c.GetHeader("Authorization")
	token, isvalid := IsValid(AuthorizationToken)
	if isvalid {
		claims, _ := token.Claims.(jwt.MapClaims)
		fetch_sql_statement := fmt.Sprintf("SELECT events.event_title, events.event_content FROM events where event_owner_name = '%s' and event_id = %s", claims["username"], post_id)
		row := db.DB.QueryRow(fetch_sql_statement)
		err_scan := row.Scan(&event.EventTitle, &event.EventContent)
		if sql.ErrNoRows == err_scan {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Event not found",
			})
		} else {
			update_sql_statement := fmt.Sprintf("UPDATE events SET event_title = '%s', event_content = '%s' where event_owner_name = '%s' and event_id = %s", jsonData.EventTitle, jsonData.EventContent, claims["username"], post_id)
			db.DB.Exec(update_sql_statement)
			c.JSON(http.StatusOK, gin.H{
				"message": "Event Updated",
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization Failed",
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
