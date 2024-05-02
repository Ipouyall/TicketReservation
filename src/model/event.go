package model

import (
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
