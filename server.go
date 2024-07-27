package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func startServer(r *chi.Mux) {
	fmt.Println("Server is running on port 4000")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
