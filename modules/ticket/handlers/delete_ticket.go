package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srgrcp/dvp-api/modules/ticket/repositories"
	"github.com/srgrcp/dvp-api/utils"
)

// DeleteTicket router endpoint of ticket module
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := mux.Vars(r)["id"]
	ticket, err := repositories.NewTicketRepository().Delete(ticketID)
	if err != nil {
		utils.Response(w, err, nil)
		return
	}
	utils.Response(w, nil, ticket)
}
