package http

import (
	"encoding/json"
	"net/http"
)

func (s *Server) setReservationHandler(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	if err != nil {
	}
}

func (s *Server) createEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) getEventsHandler(w http.ResponseWriter, r *http.Request) {

}
