package main

// import "fmt"

import (
	"TicketReservation/src/manager"
	"TicketReservation/src/rest/server"
	"log"
	"sync"
)

// func main() {
// 	fmt.Println("Server is running")
// }

func main() {
	port := "8000"

	ticketService := manager.TicketService{
		Events:  sync.Map{},
		Tickets: sync.Map{},
	}

	server := server.Server{
		TicketService: ticketService,
	}

	err := server.SetupHttpApiServer(port)
	if err != nil {
		log.Fatalf("Failed to set up HTTP server: %v", err)
	}
}
