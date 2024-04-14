package utils

import (
	"fmt"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
	}
}

func DeleteElementFromEventSlice(slice []Event, index int) []Event {
	return append(slice[:index], slice[index+1:]...)
}

func CheckForEvent(e Event) Event {
	for _, event := range events {
		if event.EventID == e.EventID {
			return event
		}
	}
	return Event{}
}
