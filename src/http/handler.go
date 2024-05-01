package http

import (
	"encoding/json"
	"net/http"
)

func (s *HttpServer) setReservationHandler(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	if err != nil {
	}
}

func (s *HttpServer) createEventHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *HttpServer) getEventsHandler(w http.ResponseWriter, r *http.Request) {

}
