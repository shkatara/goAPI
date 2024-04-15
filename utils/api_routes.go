package utils

import (
	"math/rand"
	"net/http"

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

func GetAllEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"events_list": events,
	})
}

func AddEvent(c *gin.Context) {
	event_id := rand.Intn(100000)
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	events = append(events, Event{
		EventName:  jsonData.EventName,
		EventOwner: jsonData.EventOwner,
		EventID:    event_id,
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "Event added",
	})
}

func FetchEvent(c *gin.Context) {
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	_, event := CheckForEvent(jsonData)
	if event.EventID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"Event Name":  event.EventName,
			"Event Owner": event.EventOwner,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found"},
		)
	}
}

func DeleteEvent(c *gin.Context) {
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	index, _ := CheckForEvent(jsonData)
	if index < 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found"})
	} else {
		events = DeleteElementFromEventSlice(events, index)
		c.JSON(http.StatusOK, gin.H{
			"message": "Event Deleted"})
	}
}

func UpdateEvent(c *gin.Context) {
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	CheckError(err)
	index, _ := CheckForEvent(jsonData)
	if index < 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found"})
	} else {
		events[index].EventName = jsonData.EventName
		events[index].EventOwner = jsonData.EventOwner
		c.JSON(http.StatusOK, gin.H{
			"message": events[index].EventName + " updated"})
	}
}

func GetRoot(c *gin.Context) {
	returnData := map[string]string{"name": "Shubham"}
	c.JSON(http.StatusOK, gin.H{
		"name":   returnData["name"],
		"events": listOfEvents[0]["event_name"],
	})
}
