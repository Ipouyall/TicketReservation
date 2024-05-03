package main

import (
	"TicketReservation/src/manager"
	"TicketReservation/src/rest/server"
	"log"
	"sync"
)

func main() {
	port := "8000"
	filePath := "src/app/server/database.json"

	ticketService := manager.TicketService{
		Events:  sync.Map{},
		Tickets: sync.Map{},
	}

	ticketService.ReadData(filePath)
	defer ticketService.WriteData(filePath)

	server := server.Server{
		TicketService: ticketService,
	}

	err := server.SetupHttpApiServer(port)
	if err != nil {
		log.Fatalf("Failed to set up HTTP server: %v", err)
	}
}
