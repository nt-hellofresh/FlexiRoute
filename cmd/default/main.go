package main

import (
	"github.com/nt-hellofresh/flexiroute/internal"
	"github.com/nt-hellofresh/flexiroute/pkg/flexiroute"
	"log"
)

func main() {
	log.Println("Starting server on port 5001...")
	log.Println("http://localhost:5001")

	app := createApp()

	if err := app.ServeHTTP(":5001"); err != nil {
		log.Fatalf("Failed to start server, %v", err.Error())
	}
}

func createApp() flexiroute.RouterFacade {
	router := flexiroute.NewDefaultRouter()
	internal.Configure(router)
	return router
}
