package server

import (
	"TicketReservation/src/manager"
	"TicketReservation/src/rest"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct { // we need to have access to tickets and events
	TicketService manager.TicketService
}

func (s *Server) SetupHttpApiServer(port string) error {
	router := mux.NewRouter()
	router.HandleFunc(rest.ApiSetReservation, s.setReservationHandler)
	router.HandleFunc(rest.ApiCreateEvent, s.createEventHandler)
	router.HandleFunc(rest.ApiGetEvents, s.getEventsHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         rest.ServerAddr + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
