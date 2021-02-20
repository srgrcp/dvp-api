package setup

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/srgrcp/dvp-api/modules/ticket"
	"github.com/srgrcp/dvp-api/utils"
)

//InitServer setup and serve the API
func InitServer() {
	// Automigrate database and generate test data
	err := utils.MigrateDatabase()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.Use(utils.HeadersMiddleware)
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Define module routers
	ticket.SetupTicketRouter(apiRouter)

	port := os.Getenv("PORT")
	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
