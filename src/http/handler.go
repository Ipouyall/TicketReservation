package http

import (
	"encoding/json"
	"net/http"
)

func (s *Server) setReservationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var requestData map[string]interface{}
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	
	eventID := requestData["EventID"].(string)
	numTickets := int(requestData["numTickets"].(float64))

	ticketIDs, err := s.BookingClient.BookTicket(eventID, numTickets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData := map[string]interface{}{
		"ticketIDs": ticketIDs,
	}
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (s *Server) getEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := s.BookingClient.ShowAvailableEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (s *Server) createEventHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var eventData map[string]interface{}
	err := decoder.Decode(&eventData)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	name := eventData["Name"].(string)
	data := eventData["Data"].(string)
	totalTickets := int(eventData["totalTickets"].(float64))

	event, err := s.BookingClient.CreateEvent(name, data, totalTickets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData := map[string]interface{}{
		"EventID": event.ID,
	}
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
