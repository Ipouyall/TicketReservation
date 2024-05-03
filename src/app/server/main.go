package main

import (
	"TicketReservation/src/manager"
	"TicketReservation/src/rest/server"
	"log"
	"sync"

	"os"
    "os/signal"
    "syscall"

)

func main() {
	sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	port := "8000"
	filePath := "src/app/server/database.json"

	ticketService := manager.TicketService{
		Events:  sync.Map{},
		Tickets: sync.Map{},
	}

	ticketService.ReadData(filePath)

	server := server.Server{
		TicketService: ticketService,
	}

	errCh := make(chan error, 1)

    go func() {
        errCh <- server.SetupHttpApiServer(port)
    }()

	select {
    case sig := <-sigCh:
        log.Printf("Received signal: %v", sig)
		server.TicketService.WriteData(filePath)
    case err := <-errCh:
        if err != nil {
            log.Fatalf("Failed to set up HTTP server: %v", err)
        }
    }
}
