package flexiroute

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func NewChiHandler(opts ...RouteSpecOpts) http.Handler {
	rs := NewRouteSpec(opts...)
	app := chi.NewRouter()
	buildChiRoutes(rs, app)
	return app
}

func buildChiRoutes(rs *RouteSpecification, router chi.Router) {
	for _, route := range rs.routes {
		for _, mw := range rs.middlewares {
			route.WithMiddleWare(mw)
		}

		registerEndpoint(rs, router, route)
	}

	for _, ns := range rs.subRoutes {
		buildChiRoutes(ns, router)
	}
}

func registerEndpoint(rs *RouteSpecification, app chi.Router, route *ApiRoute) {
	path := route.Path()

	if rs.name != "" {
		// Slightly cheating here as proper nested chi routes aren't
		// really utilised. Instead, the fully qualified path is
		// registered with the application.
		path = fmt.Sprintf("/%v%v", rs.name, route.Path())
	}

	handler := route.ToHandlerFunc(rs.templates)
	switch route.Method() {
	case http.MethodGet:
		app.Get(path, handler)
	case http.MethodPut:
		app.Put(path, handler)
	case http.MethodPatch:
		app.Patch(path, handler)
	case http.MethodPost:
		app.Post(path, handler)
	case http.MethodDelete:
		app.Delete(path, handler)
	default:
		log.Fatalf("unsupported method")
	}
}
