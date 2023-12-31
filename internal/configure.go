package internal

import (
	"github.com/nt-hellofresh/flexiroute/internal/middleware"
	"github.com/nt-hellofresh/flexiroute/internal/routes"
	"github.com/nt-hellofresh/flexiroute/pkg/flexiroute"
)

func Specification() []flexiroute.RouteSpecOpts {
	return []flexiroute.RouteSpecOpts{
		flexiroute.HtmlTemplatesDir("internal/www/*.html"),
		flexiroute.Use(middleware.LoggingMiddleware),
		flexiroute.Get("/", routes.HomeHandler),
		flexiroute.Get("/dice", routes.RandomNumberHandler),
		flexiroute.Namespace("users",
			flexiroute.Get("/", routes.GetUsersHandler),
			flexiroute.Put("/test", routes.PutHandler),
		),
	}
}
