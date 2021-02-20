package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srgrcp/dvp-api/modules/ticket/repositories"
	"github.com/srgrcp/dvp-api/utils"
)

// GetTicket router endpoint of ticket module
func GetTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := mux.Vars(r)["id"]
	ticket, err := repositories.NewTicketRepository().Find(ticketID)
	if err != nil {
		utils.Response(w, err, nil)
		return
	}
	utils.Response(w, nil, ticket)
}
