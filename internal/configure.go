package internal

import (
	"github.com/nt-hellofresh/flexiroute/internal/middleware"
	"github.com/nt-hellofresh/flexiroute/internal/routes"
	"github.com/nt-hellofresh/flexiroute/pkg/flexiroute"
)

func Configure(router flexiroute.RouterFacade) {
	router.Use(middleware.LoggingMiddleware)

	router.Get("/", routes.HomeHandler)
	router.Get("/dice", routes.RandomNumberHandler)

	users := router.Namespace("users")
	users.Get("/", routes.GetUsersHandler)
	users.Put("/test", routes.PutHandler)

	router.LoadTemplates("internal/www/*.html")
}
