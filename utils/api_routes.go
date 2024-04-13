package utils

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

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

func SaveEvent(c *gin.Context) {
	event_id := rand.Intn(100000)
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	if err != nil {
		fmt.Println("Error in binding JSON: ", err)
	}
	events = append(events, Event{
		EventName:  jsonData.EventName,
		EventOwner: jsonData.EventOwner,
		EventID:    event_id,
	})

}

func GetRoot(c *gin.Context) {
	returnData := map[string]string{"name": "Shubham"}
	c.JSON(200, gin.H{
		"name": returnData["name"],
	})
}
