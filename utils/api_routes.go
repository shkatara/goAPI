package utils

import (
	"database/sql"
	"net/http"

	"example.com/api/db"
	"github.com/gin-gonic/gin"
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

var events = []Event{
	{EventName: "A Go-lang life", EventOwner: "Shubham", EventID: 1},
	{EventName: "There goes another", EventOwner: "Shubham", EventID: 2},
}

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}

func GetAllEvents(c *gin.Context) {
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
