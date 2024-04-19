package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./db.sqlite3")

	if err != nil {
		fmt.Println(err)
		panic("Connection to database failed")
	} else {
		fmt.Println("Database opened successfully")
	}
}

func CreateEventTable() {
	eventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_title VARCHAR(255) NOT NULL,
		event_owner VARCHAR(255) NOT NULL
	)
	`
	_, err := DB.Exec(eventTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create events table")
	} else {
		fmt.Println("Created events table successfully")
	}

}
