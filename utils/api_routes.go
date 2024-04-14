package utils

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
	}
}

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
	c.JSON(200, gin.H{
		"events_list": events,
	})
}

func AddEvent(c *gin.Context) {
	event_id := rand.Intn(100000)
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	checkError(err)
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
	checkError(err)
	for _, event := range events {
		if event.EventID == jsonData.EventID {
			c.JSON(http.StatusOK, gin.H{
				"event_name":  event.EventName,
				"event_owner": event.EventOwner,
			})
			return // This is important to stop the loop if one hit is found.
			// Need to stop the loop if the event is found as event_id is unique
			// Otherwise, it will keep on searching.
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Event not found",
	})
}

func deleteElementFromEventSlice(slice []Event, index int) []Event {
	return append(slice[:index], slice[index+1:]...)
}

func DeleteEvent(c *gin.Context) {
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	checkError(err)
	for i, event := range events {
		if event.EventID == jsonData.EventID {
			fmt.Println("Event is", event.EventName, "available at index", i)
			events = deleteElementFromEventSlice(events, event.EventID)
			c.JSON(200, gin.H{
				"message": "Event deleted",
			})
		}
	}

}

func GetRoot(c *gin.Context) {
	returnData := map[string]string{"name": "Shubham"}
	c.JSON(200, gin.H{
		"name":   returnData["name"],
		"events": listOfEvents[0]["event_name"],
	})
}
