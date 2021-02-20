package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srgrcp/dvp-api/modules/ticket/models"
	"github.com/srgrcp/dvp-api/modules/ticket/repositories"
	"github.com/srgrcp/dvp-api/utils"
)

// UpdateTicket router endpoint of ticket module
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := mux.Vars(r)["id"]
	ticketInput := &models.UpdateTicketInput{}
	json.NewDecoder(r.Body).Decode(ticketInput)

	ticket, err := repositories.NewTicketRepository().Update(ticketID, ticketInput)
	if err != nil {
		utils.Response(w, err, nil)
		return
	}
	utils.Response(w, nil, ticket)
}
