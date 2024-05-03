package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (s *Server) setReservationHandler(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("[Server] (ticket reservation)")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var requestData map[string]interface{}
	err := decoder.Decode(&requestData)
	if err != nil {
		log.Println("Failed to parse request body.")
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	eventID := requestData["EventID"].(string)
	numTickets := int(requestData["numTickets"].(float64))

	ticketIDs, err := s.TicketService.BookTickets(eventID, numTickets)
	if err != nil {
		log.Println("Failed to book tickets:", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData := map[string]interface{}{
		"ticketIDs": ticketIDs,
	}
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		log.Println("Failed to encode response:", err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	log.Println("Reservation successful.")
}

func (s *Server) getEventsHandler(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("[Server] (get events)")
	events := s.TicketService.ListEvents()

	err := json.NewEncoder(w).Encode(events)
	if err != nil {
		log.Println("Failed to encode response:", err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	log.Println("Events retrieved successfully.")
}

func (s *Server) createEventHandler(w http.ResponseWriter, r *http.Request) {
	log.SetPrefix("[Server] (create event)")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var eventData map[string]interface{}
	err := decoder.Decode(&eventData)
	if err != nil {
		log.Println("Failed to parse request body:", err.Error())
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	fmt.Println(eventData)

	name := eventData["Name"].(string)
	date, err := time.Parse("2006-01-02 15:04", eventData["Date"].(string))
	if err != nil {
		log.Println("Failed to parse date:", err.Error())
		http.Error(w, "Failed to parse date", http.StatusBadRequest)
		return
	}
	totalTickets := int(eventData["totalTickets"].(float64))

	event, err := s.TicketService.CreateEvent(name, date, totalTickets)
	if err != nil {
		log.Println("Failed to create event:", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData := map[string]interface{}{
		"EventID": event.ID,
	}
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		log.Println("Failed to encode response:", err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	log.Println("Event created successfully: ", event)
}
