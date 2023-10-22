package flexiroute

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

func BuildChiHandler(opts ...RouteSpecOpts) http.Handler {
	rs := NewRouteSpec(opts...)
	app := chi.NewRouter()
	buildChiRoutes(rs, app)
	return app
}

func buildChiRoutes(rs *RouteSpecification, router chi.Router) {
	for _, mw := range rs.middlewares {
		router.Use(mw.ToHandler())
	}
	for _, route := range rs.routes {
		registerEndpoint(rs.templates, router, route)
	}

	for _, ns := range rs.subRoutes {
		router.Group(func(r chi.Router) {
			r.Route(fmt.Sprintf("/%v", ns.name), func(r chi.Router) {
				buildChiRoutes(ns, r)
			})
		})
	}
}

func registerEndpoint(templates *template.Template, app chi.Router, route *ApiRoute) {
	path := route.Path()
	handler := route.ToHandlerFunc(templates)

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
