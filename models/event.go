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
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		fmt.Println("error in executing the statement")
		return err
	}
	return err
}

func (e Event) print() {
	fmt.Println(e.toString())
}

func (e Event) toString() string {
	return "Name:" + e.Name + ",Desc:" + e.Description
}

func GetAllEvents() []Event {
	events = []Event{} // Clear the events slice to avoid duplication
	query := `SELECT id, name, description, location, dateTime , user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("error in getting the events")
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		// Print raw row values before scanning (for debugging)
		cols, err := rows.Columns()
		if err != nil {
			fmt.Println("error getting columns:", err)
			continue
		}
		rawResult := make([]interface{}, len(cols))
		dest := make([]interface{}, len(cols))
		for i := range rawResult {
			dest[i] = &rawResult[i]
		}
		if err := rows.Scan(dest...); err != nil {
			fmt.Println("error scanning raw row:", err)
			continue
		}
		fmt.Println("Raw row values:", rawResult)

		// Now scan into the event struct as usual
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			fmt.Println("error in scanning the event", err)
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
	query := `DELETE FROM events WHERE id=? or id is null`
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
	query := `SELECT id, name, description, location, dateTime,user_id FROM events WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("error in preparing the statement")
		return Event{}
	}
	defer stmt.Close()
	var event Event
	err = stmt.QueryRow(id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		fmt.Println("error in scanning the event ," + err.Error())
		return Event{}
	}
	return event
}
