package manager

import (
	"TicketReservation/src/model"
	"TicketReservation/src/utils"
	"fmt"
	"sync"
	"time"
)

type TicketService struct {
	Events      sync.Map // map[int]*model.Event
	Tickets     sync.Map // map[int]*model.Ticket
	EventMutex  sync.Mutex
	TicketMutex sync.Mutex
}

func (ts *TicketService) CreateEvent(name string, data time.Time, totalTickets int) (*model.Event, error) {

	event := model.NewEvent(utils.GenerateUUID(), name, data, totalTickets, totalTickets)

	ts.Events.Store(event.ID, event)
	return &event, nil
}

func (ts *TicketService) ListEvents() []model.Event {
	var events []model.Event
	ts.Events.Range(func(key, value interface{}) bool {
		event := value.(model.Event)
		events = append(events, event)
		return true
	})
	return events
}

func (ts *TicketService) BookTickets(eventID string, numTickets int) ([]string, error) {

	//implement concurrency control here (step 3)

	ts.EventMutex.Lock()
	defer ts.EventMutex.Unlock()

	eventObj, ok := ts.Events.Load(eventID)
	if !ok {
		return nil, fmt.Errorf("model not found")

	}

	ev := eventObj.(model.Event)
	if ev.AvailableTickets < numTickets {
		return nil, fmt.Errorf("not enough Tickets available")
	}

	var ticketIDs []string
	for i := 0; i < numTickets; i++ {
		ticketID := utils.GenerateUUID()
		ticketIDs = append(ticketIDs, ticketID)
		// Store the ticket in a separate model structure if needed
	}

	ev.AvailableTickets -= numTickets
	ts.Events.Store(eventID, ev)

	return ticketIDs, nil

}
