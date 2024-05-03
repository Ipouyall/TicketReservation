package rest

import (
	"TicketReservation/src/manager"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct { // we need to have access to tickets and events
	TicketService manager.TicketService
}

func (s *Server) SetupHttpApiServer(port string) error {
	router := mux.NewRouter()
	router.HandleFunc(apiSetReservation, s.setReservationHandler)
	router.HandleFunc(apiCreateEvent, s.createEventHandler)
	router.HandleFunc(apiGetEvents, s.getEventsHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         serverAddr + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
