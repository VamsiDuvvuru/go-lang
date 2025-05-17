package models

import (
	"fmt"
	"time"
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

func (e Event) Save() {
	events = append(events, e)
}

func (e Event) print() {
	fmt.Println(e.toString())
}

func (e Event) toString() string {
	return "Name:" + e.Name + ",Desc:" + e.Description
}

func GetAllEvents() []Event {
	return events
}

func UpdateEvent(id int, event Event) {
	for i, val := range events {
		fmt.Println("inside updating the event")
		fmt.Println("event id is", val.ID)
		if events[i].ID == id {
			tempEvent := &events[i]
			tempEvent.update(event)
			return
		}
	}
}

func (event *Event) update(newEvent Event) {
	fmt.Println("inside updating the new event")
	event.ID = newEvent.ID
	event.DateTime = newEvent.DateTime
	event.Description = newEvent.Description
	event.Location = newEvent.Location
	event.Name = newEvent.Name
	fmt.Println("event is updated successfully")
}
