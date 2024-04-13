package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Event struct {
	EventName  string `json:"event_name" binding:"required"`
	EventOwner string `json:"event_owner" binding:"required"`
}

var events = []Event{
	{EventName: "A Go-lang life", EventOwner: "Shubham"},
	{EventName: "There goes another", EventOwner: "Shubham"},
}

func GetAllEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"events": events,
	})
}

func SaveEvent(c *gin.Context) {
	var jsonData Event
	err := c.ShouldBindJSON(&jsonData)
	if err != nil {
		fmt.Println("Error in binding JSON: ", err)
	}
	fmt.Println("Event Name: ", jsonData)
	events = append(events, Event{
		EventName: jsonData.EventName, EventOwner: jsonData.EventOwner,
	})

}

func GetRoot(c *gin.Context) {
	returnData := map[string]string{"name": "Shubham"}
	c.JSON(200, gin.H{
		"name": returnData["name"],
	})
}
