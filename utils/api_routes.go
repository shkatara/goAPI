package utils

import (
	"fmt"
	"math/rand"

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

var events = []Event{
	{EventName: "A Go-lang life", EventOwner: "Shubham", EventID: 1},
	{EventName: "There goes another", EventOwner: "Shubham", EventID: 2},
}

func GetAllEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"events": events,
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

}

func FetchEvent(c *gin.Context) {
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	checkError(err)
	for _, event := range events {
		if event.EventID == jsonData.EventID {
			c.JSON(200, gin.H{
				"event": event,
			})
		}
	}
}

func GetRoot(c *gin.Context) {
	returnData := map[string]string{"name": "Shubham"}
	c.JSON(200, gin.H{
		"name": returnData["name"],
	})
}
