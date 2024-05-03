package model

import (
	"fmt"
	"time"
)

type Event struct {
	ID               string
	Name             string
	Date             time.Time
	TotalTickets     int
	AvailableTickets int
}

func NewEvent(id, name string, date time.Time, totalTickets, AvailableTickets int) Event {
	return Event{
		ID:               id,
		Name:             name,
		Date:             date,
		TotalTickets:     totalTickets,
		AvailableTickets: AvailableTickets,
	}
}

func (event Event) ToString() string {
	return fmt.Sprintf("%v\n\tName: %v\n\tDate: %v\n\tAvailable Tickets: %v of %v",
		event.ID, event.Name, event.Date.Format("2024-01-01 11:11"), event.AvailableTickets, event.TotalTickets)
}
