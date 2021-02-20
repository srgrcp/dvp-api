package handlers

import (
	"net/http"

	"github.com/srgrcp/dvp-api/modules/ticket/repositories"
	"github.com/srgrcp/dvp-api/utils"
)

// GetTickets router endpoint of ticket module
func GetTickets(w http.ResponseWriter, r *http.Request) {
	page := repositories.NewTicketRepository().Query(r)
	utils.Response(w, nil, page)
}
