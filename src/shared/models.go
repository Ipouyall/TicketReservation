package shared

import (
	"fmt"
	"sync"
	"time"
)

type event struct {
	ID               string
	Name             string
	Date             time.Time
	TotalTickets     int
	AvailableTickets int
}

type ticket struct {
	ID      string
	EventID string
}

type TicketService struct {
	events      sync.Map
	tickets     map[int]*ticket
	eventMutex  sync.Mutex
	ticketMutex sync.Mutex
}

func (ts *TicketService) CreateEvent(name string, data time.Time, totalTickets int) (*event, error) {

	event := &event{
		ID:               generateUUID(), //Generate a unique ID for the event
		Name:             name,
		Date:             data,
		TotalTickets:     totalTickets,
		AvailableTickets: totalTickets,
	}

	ts.events.Store(event.ID, event)
	return event, nil
}

func (ts *TicketService) ListEvents() []*event {
	var events []*event
	ts.events.Range(func(key, value interface{}) bool {
		event := value.(*event)
		events = append(events, event)
		return true
	})
	return events
}

func (ts *TicketService) BookTickets(eventID string, numTickets int) ([]string, error) {

	//implement concurrency control here (step 3)

	ts.eventMutex.Lock()
	defer ts.eventMutex.Unlock()

	eventObj, ok := ts.events.Load(eventID)
	if !ok {
		return nil, fmt.Errorf("event not found")

	}

	ev := eventObj.(*event)
	if ev.AvailableTickets < numTickets {
		return nil, fmt.Errorf("not enough tickets available")
	}

	var ticketIDs []string
	for i := 0; i < numTickets; i++ {
		ticketID := generateUUID()
		ticketIDs = append(ticketIDs, ticketID)
		// Store the ticket in a separate data structure if needed
	}

	ev.AvailableTickets -= numTickets
	ts.events.Store(eventID, ev)

	return ticketIDs, nil

}
