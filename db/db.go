package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("could not connect to db: " + err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateEventTable()
}

func CreateEventTable() {
	query := `
		CREATE TABLE IF NOT EXISTS EVENTS (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER NOT NULL
		);
    `
	_, err := DB.Exec(query)
	if err != nil {
		panic("could not create events table: " + err.Error())
	}
}
