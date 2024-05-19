package events

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
	}
}

func DeleteElementFromEventSlice(slice []Event, index int) []Event {
	return append(slice[:index], slice[index+1:]...)
}

func CheckForEvent(e Event) (int, Event) {
	for i, event := range events {
		if event.EventID == e.EventID {
			return i, event
		}
	}
	return -1, Event{}
}

func IsValid(AuthorizationToken string) (*jwt.Token, bool) {
	token, err := jwt.Parse(AuthorizationToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err == nil {
		return token, true
	} else {
		return nil, false
	}
}
