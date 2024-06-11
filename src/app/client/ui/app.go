package ui

import (
	"TicketReservation/src/rest"
	"TicketReservation/src/rest/client"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	api client.IClient
}

func NewApp(api client.IClient) *App {
	return &App{api: api}
}

func (a *App) showEvents() {
	events, _, err := a.api.GetEvents()
	if err != nil {
		fmt.Println("Failed to get events from server.")
		return
	}

	for i, event := range events {
		fmt.Printf("(%v) %v\n", i+1, event.ToString())
	}
	return
}

func (a *App) bookTickets() {
	events, _, err := a.api.GetEvents()
	if err != nil {
		fmt.Println("Failed to get events from server.")
	}

	board := NewEventModel(events)
	p := tea.NewProgram(board)
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	selectedEvent := events[m.(EventModel).Selected]

	clearTerminal()
	fmt.Println("Booking Ticket for: ", selectedEvent.ToString()+"\n")

	// Get quantity of tickets
	var quantity int
	fmt.Print("Quantity: ")
	_, err = fmt.Scanln(&quantity)
	if err != nil {
		log.Fatal(err)
	}

	ticketsID, _, err := a.api.BookTicket(selectedEvent.ID, quantity)
	if err != nil {
		fmt.Println("Failed to book ticket(s)")
		return
	}

	fmt.Println("\nTickets booked successfully!\n\"Ticket ID:")
	for i, id := range ticketsID {
		fmt.Println("(", i+1, ") ", id)
	}
	return
}

func (a *App) createNewEvent() {
	var name string
	var date time.Time
	var dateStr, timeStr string
	var totalTickets int

	fmt.Println("**[Create new event]**")
	fmt.Print("Name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Date (YYYY-MM-DD hh:mm): ")
	_, err = fmt.Scanln(&dateStr, &timeStr)
	if err != nil {
		log.Fatal(err)
	}
	date, err = time.Parse("2006-01-02 15:04", dateStr+" "+timeStr)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Total tickets: ")
	_, err = fmt.Scanln(&totalTickets)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	id, _, err := a.api.CreateEvent(name, date, totalTickets)

	if err != nil {
		fmt.Printf("Failed to create event.")
		return
	}
	fmt.Println("Event created successfully!\n\"Event ID: ", id, "\"")
	return
}

func (a *App) help() {
	fmt.Println("**[Help menu]**")

	fmt.Println("\nYour are in the interactive UI, here are what you can do:")
	fmt.Println("1. Show events: You can see are events defined in the system, Events won't be filtered by date for ticket availability and sorted in random order.")
	fmt.Println("2. Book ticket: You can book tickets for an event if enough number of ticket is remained.")
	fmt.Println("3. Create new event: You can create a new event, you can define the name, date and total tickets for it.")
	fmt.Println("4. Help: You can see this menu again.")
	fmt.Println("5. Exit: You can exit the program. You can also use \"ctrl+c\", \"esc\", \"q\" to exit.")

	fmt.Println("\nThere also exists a test mode, designed to test server under pressure. You can't access that mode here.")
	fmt.Println("To enable test mode, run the program with the -test -client <number of parallel clients> -pressure <number of request each client send in each test stage>")
	fmt.Println("\tExample: ./client -test -client 5 -pressure 10")
	fmt.Println("\tThis will run 5 parallel clients, each sending 10 requests in each test stage.")

	fmt.Println("\nTo run client, use -p flag to set port. default is ", rest.DefaultPort, ".")
	return
}

func (a *App) step() bool {
	initialModel := Model{}
	initialModel.InitBaseMenu()

	p := tea.NewProgram(initialModel)
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}

	idx := m.(Model).Selected
	clearTerminal()
	if idx == -1 {
		return false
	}
	itemID := initialModel.Items[idx].ID
	switch itemID {
	case 1: // Show Events
		a.showEvents()
	case 2: // Book Ticket
		a.bookTickets()
	case 3: // New event
		a.createNewEvent()
	case 4: // Help
		a.help()
	case 5: // Exit
		return false
	}
	return true
}

func (a *App) RunUI() {
	for {
		clearTerminal()
		if a.step() == false {
			fmt.Print("\033[2J")
			fmt.Println("Exiting the program...")
			return
		}
		// Wait for used to press enter to return
		fmt.Println("\nPress enter to return")
		_, _ = fmt.Scanln()

		time.Sleep(100 * time.Millisecond)
	}
}

func clearTerminal() {
	fmt.Print("\033[2J")
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) RunTest(clientNum, loadDeg int) {
	var wg sync.WaitGroup

	for i := 0; i < clientNum; i++ {
		wg.Add(1)

		go func(clientID int) {
			defer wg.Done()
			for j := 0; j < loadDeg; j++ {
				a.showEvents()
			}
			fmt.Printf("Client %d completed all requests.\n", clientID)
		}(i)
	}

	wg.Wait()
	fmt.Println("All clients have finished.")
}
