package manager

import (
	"TicketReservation/src/model"
	"TicketReservation/src/utils"
	"fmt"
	"sync"
	"time"
	"encoding/json"
	"io/ioutil"
	"os"
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

func (ts *TicketService) ReadData(filePath string) {
	ts.EventMutex.Lock()
	defer ts.EventMutex.Unlock()
	
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	var events []model.Event
	err = json.Unmarshal(data, &events)
	if err != nil {
		return
	}

	for _, event := range events {
		ts.Events.Store(event.ID, event)
	}

	return
}

func (ts *TicketService) WriteData(filePath string) {
	ts.EventMutex.Lock()
	defer ts.EventMutex.Unlock()

	fmt.Println(ts.ListEvents())

	eventsJSON, err := json.Marshal(ts.ListEvents())
	if err != nil {
		return
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	err = ioutil.WriteFile(filePath, eventsJSON, 0644)
	if err != nil {
		return
	}

	return
}