package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"TicketReservation/src/app/server/manager"
)

type Server struct { // we need to have access to tickets and events
	ticketservice: manager.TicketService,
}

func (s *Server) SetupHttpApiServer() error {
	router := mux.NewRouter()
	router.HandleFunc(apiSetReservation, s.setReservationHandler)
	router.HandleFunc(apiCreateEvent, s.createEventHandler)
	router.HandleFunc(apiGetEvents, s.getEventsHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         serverAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
