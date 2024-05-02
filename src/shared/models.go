
type event struct {
    ID      string
    Name    string
    Date    time.Time
	TotalTickets int
    AvailableTickets  int
}

type ticket struct {
    ID     string
    EventID  string
}

type TicketService struct {
    events     sync.Map
    tickets    map[int]*ticket
    eventMutex sync.Mutex
    ticketMutex sync.Mutex
}

