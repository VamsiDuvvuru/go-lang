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

func (e Event) showEvents() {
	for _, j := range events {
		j.print()
	}
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
