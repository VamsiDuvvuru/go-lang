package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	fmt.Println("Initializing DB")
	DB, err := sql.Open("postgres",
		"host=localhost port=5432 user=postgres password=talent dbname=Events sslmode=disable")
	if err != nil {
		fmt.Println("DB open error:", err)
	}
	if err = DB.Ping(); err != nil {
		fmt.Println("DB ping error:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	fmt.Println("DB initialized")
	CreateEventTable()
}

func CreateEventTable() {
	query := `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime TIMESTAMP NOT NULL,
        user_id INTEGER NOT NULL
    );
    `
	_, err := DB.Exec(query)
	if err != nil {
		panic("could not create events table: " + err.Error())
	}
}
