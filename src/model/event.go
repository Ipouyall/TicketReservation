package model

import (
	"fmt"
	"time"
)

type Event struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Date             time.Time `json:"date"`
	TotalTickets     int       `json:"totalTickets"`
	AvailableTickets int       `json:"availableTickets"`
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
		event.ID, event.Name, event.Date.Format("2006-01-02 15:04"), event.AvailableTickets, event.TotalTickets)
}
