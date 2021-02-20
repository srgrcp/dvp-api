package ticket

import (
	"github.com/gorilla/mux"
	"github.com/srgrcp/dvp-api/modules/ticket/handlers"
)

// SetupTicketRouter ticket's module's router
func SetupTicketRouter(apiRouter *mux.Router) *mux.Router {
	ticketRouter := apiRouter.PathPrefix("/ticket").Subrouter()

	ticketRouter.HandleFunc("/", handlers.CreateTicket).Methods("POST")
	ticketRouter.HandleFunc("/q", handlers.GetTickets).Methods("GET")
	ticketRouter.HandleFunc("/{id}", handlers.GetTicket).Methods("GET")
	ticketRouter.HandleFunc("/{id}", handlers.UpdateTicket).Methods("PUT")
	ticketRouter.HandleFunc("/{id}", handlers.DeleteTicket).Methods("DELETE")

	return ticketRouter
}
