package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/srgrcp/dvp-api/modules/ticket/models"
	"github.com/srgrcp/dvp-api/modules/ticket/repositories"
	"github.com/srgrcp/dvp-api/utils"
)

// CreateTicket router endpoint of ticket module
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	ticketInput := &models.CreateTicketInput{}
	err := json.NewDecoder(r.Body).Decode(ticketInput)
	if err != nil {
		panic(err)
		//utils.Response(w, err, nil)
		//return
	}

	ticket, err := repositories.NewTicketRepository().Create(ticketInput)
	if err != nil {
		utils.Response(w, err, nil)
		return
	}
	utils.Response(w, nil, ticket)
}
