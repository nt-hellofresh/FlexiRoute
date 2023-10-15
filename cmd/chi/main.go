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

	app := createApp()

	if err := http.ListenAndServe(":5001", app); err != nil {
		log.Fatalf("Failed to start server, %v", err.Error())
	}
}

func createApp() flexiroute.RouterFacade {
	router := flexiroute.NewChiRouter()
	internal.Configure(router)
	return router
}
