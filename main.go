package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"investmentTracker/api"
	"net/http"
)

// GO MUX Documentation: https://github.com/gorilla/mux

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/report", api.ReportInvestment).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server is listening on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error:", err)
	}
}
