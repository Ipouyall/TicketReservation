package client

import (
	"TicketReservation/src/app/client/ui"
	"TicketReservation/src/rest"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"os/exec"
	"time"
)

type App struct {
	api rest.IClient
}

func NewApp(api rest.IClient) *App {
	return &App{api: api}
}

func (a *App) showEvents() {
	events, _, err := a.api.GetEvents()
	if err != nil {
		fmt.Printf("Error getting events: %v\n", err)
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
		fmt.Printf("Error getting events: %v\n", err)
	}

	board := ui.NewEventModel(events)
	p := tea.NewProgram(board)
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	selectedEvent := events[m.(ui.EventModel).Selected]

	fmt.Print("\033[2J")
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
		fmt.Printf("Error booking ticket: %v\n", err)
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
	var totalTickets int

	fmt.Println("**[Create new event]**")
	fmt.Print("Name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Date (YYYY-MM-DD hh:mm): ")
	_, err = fmt.Scanln(&date)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Total tickets: ")
	_, err = fmt.Scanln(&totalTickets)
	if err != nil {
		log.Fatal(err)
	}

	id, _, err := a.api.CreateEvent(name, date, totalTickets)

	if err != nil {
		fmt.Printf("Error creating event: %v\n", err)
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
	fmt.Println("5. Exit: You can exit the program.")

	fmt.Println("\nThere also exists a test mode, designed to test server inder pressure. You can't access that mode here.")
	fmt.Println("To enable test mode, run the program with the -test -client <number of parallel clients> -pressure <number of request each client send in each test stage>")
	fmt.Println("\tExample: ./client -test -client 5 -pressure 10")
	fmt.Println("\tThis will run 5 parallel clients, each sending 10 requests in each test stage.")

	fmt.Println("\nTo run client, use -p flag to set port. default is 8000")
	return
}

func (a *App) step() bool {
	initialModel := ui.Model{}
	initialModel.InitBaseMenu()

	p := tea.NewProgram(initialModel)
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}

	itemID := initialModel.Items[m.(ui.Model).Selected].ID
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

func (a *App) Run() {
	for {
		fmt.Print("\033[2J")
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		if a.step() == false {
			break
		}
		// Wait for used to press enter to return
		fmt.Println("\nPress enter to return")
		_, _ = fmt.Scanln()

		time.Sleep(1 * time.Second)
	}
}
