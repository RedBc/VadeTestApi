package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"vade.com/vade/goDocumentAPI/handlers"
)

func main() {
	// Create new Router
	router := mux.NewRouter()

	// route properly to respective handlers
	router.Handle("/documents", handlers.GetDocumentsHandler()).Methods("GET")

	// Create new server and assign the router
	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	fmt.Println("Staring Vade Api server on Port 9090")
	// Start Server on defined port/host.
	server.ListenAndServe()
}
