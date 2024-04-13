package utils

import (
	"github.com/gin-gonic/gin"
)

type Event struct {
	EventName  string `json:"event_name"`
	EventOwner string `json:"event_date"`
}

var events = []Event{
	{EventName: "A Go-lang life", EventOwner: "Shubham"},
	{EventName: "There goes another", EventOwner: "Shubham"},
}

func (e Event) AddEvent() {
	events = append(events, Event{
		EventName: "A Go-lang life", EventOwner: "Shubham",
	},
	)
}

func GetAllEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"events": events,
	})
}

func GetRoot(c *gin.Context) {
	returnData := map[string]string{"name": "Shubham"}
	c.JSON(200, gin.H{
		"name": returnData["name"],
	})
}
