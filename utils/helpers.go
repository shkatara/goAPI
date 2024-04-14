package utils

import "fmt"

func CheckError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
	}
}

func DeleteElementFromEventSlice(slice []Event, index int) []Event {
	return append(slice[:index], slice[index+1:]...)
}
