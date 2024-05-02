package manager

import (
	"TicketReservation/src/model"
	"TicketReservation/src/utils"
	"fmt"
	"sync"
	"time"
)

type TicketService struct {
	events      sync.Map
	tickets     map[int]*model.Ticket
	eventMutex  sync.Mutex
	ticketMutex sync.Mutex
}

func (ts *TicketService) CreateEvent(name string, data time.Time, totalTickets int) (*model.Event, error) {

	event := model.NewEvent(utils.GenerateUUID(), name, data, totalTickets, totalTickets)

	ts.events.Store(event.ID, event)
	return &event, nil
}

func (ts *TicketService) ListEvents() []*model.Event {
	var events []*model.Event
	ts.events.Range(func(key, value interface{}) bool {
		event := value.(*model.Event)
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
		return nil, fmt.Errorf("model not found")

	}

	ev := eventObj.(*model.Event)
	if ev.AvailableTickets < numTickets {
		return nil, fmt.Errorf("not enough tickets available")
	}

	var ticketIDs []string
	for i := 0; i < numTickets; i++ {
		ticketID := utils.GenerateUUID()
		ticketIDs = append(ticketIDs, ticketID)
		// Store the ticket in a separate model structure if needed
	}

	ev.AvailableTickets -= numTickets
	ts.events.Store(eventID, ev)

	return ticketIDs, nil

}
