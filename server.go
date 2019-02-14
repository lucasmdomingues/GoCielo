package main

import (
	"fmt"
	"gocielo/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)

	// Sandbox
	sandbox := r.PathPrefix("/sandbox").Subrouter()
	sandbox.HandleFunc("/creditcard", handlers.CreditCardHandlerSandBox)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
