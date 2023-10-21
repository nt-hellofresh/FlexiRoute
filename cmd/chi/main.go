package main

import (
	"github.com/nt-hellofresh/flexiroute/internal"
	"github.com/nt-hellofresh/flexiroute/pkg/flexiroute"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port 5001...")
	log.Println("http://localhost:5001")

	handler := flexiroute.BuildChiHandler(internal.Specification()...)

	if err := http.ListenAndServe(":5001", handler); err != nil {
		log.Fatalf("Failed to start server, %v", err.Error())
	}
}
