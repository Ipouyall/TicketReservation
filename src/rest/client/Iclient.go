package client

import (
	"TicketReservation/src/model"
	"time"
)

type IClient interface {
	GetEvents() ([]model.Event, string, error)
	BookTicket(eventId string, quantity int) ([]string, string, error)
	CreateEvent(name string, date time.Time, totalTickets int) (string, string, error)

	BurstTest() (string, error)
}
