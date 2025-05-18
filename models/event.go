package models

import (
	"fmt"
	"time"

	"example.com/mygolangproj/db"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query :=
		`INSERT INTO events(name, description, location, dateTime, user_id) 
	 VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("error in preparing the statement")
		return err
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		fmt.Println("error in executing the statement")
		return err
	}
	id, err := result.LastInsertId()
	e.ID = int(id)
	return err
}

func (e Event) print() {
	fmt.Println(e.toString())
}

func (e Event) toString() string {
	return "Name:" + e.Name + ",Desc:" + e.Description
}

func GetAllEvents() []Event {
	query := `SELECT id, name, description, location, dateTime FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("error in getting the events")
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime)
		if err != nil {
			fmt.Println("error in scanning the event")
			return nil
		}
		events = append(events, event)
	}
	return events
}

func UpdateEvent(id int, event Event) {
	query := `UPDATE events SET name=?, description=?, location=?, dateTime=? WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("error in preparing the statement")
		return
	}
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, id)
	if err != nil {
		fmt.Println("error in executing the statement")
		return
	}
	result.LastInsertId()
	fmt.Println("event updated successfully")
}

// func (event *Event) update(newEvent Event) {
// 	fmt.Println("inside updating the new event")
// 	event.ID = newEvent.ID
// 	event.DateTime = newEvent.DateTime
// 	event.Description = newEvent.Description
// 	event.Location = newEvent.Location
// 	event.Name = newEvent.Name
// 	fmt.Println("event is updated successfully")
// }

func DeleteEvent(id string) {
	query := `DELETE FROM events WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("error in preparing the statement")
		return
	}
	result, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("error in executing the statement")
		return
	}
	result.LastInsertId()
	fmt.Println("event deleted successfully")
}

func GetEventsById(id string) Event {
	query := `SELECT id, name, description, location, dateTime FROM events WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("error in preparing the statement")
		return Event{}
	}
	defer stmt.Close()
	var event Event
	err = stmt.QueryRow(id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime)
	if err != nil {
		fmt.Println("error in scanning the event")
		return Event{}
	}
	return event
}
