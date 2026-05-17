package main

import (
	"fmt"
	"jeanleeb/sudoku-solver/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	svc := &handlers.Service{}

	handlers.Register(mux, svc)

	addr := ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	fmt.Printf("Starting server on %s\n", addr)
	log.Fatal(server.ListenAndServe())
}
