package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HttpServer struct { // we need to have access to tickets and events

}

func (s *HttpServer) SetupHttpApiServer() {
	router := mux.NewRouter()
	router.HandleFunc(apiSetReservation, s.setReservationHandler)
	router.HandleFunc(apiCreateEvent, s.createEventHandler)
	router.HandleFunc(apiGetEvents, s.getEventsHandler)

	srv := &http.Server{
		Handler: router,
		//Addr: http.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
