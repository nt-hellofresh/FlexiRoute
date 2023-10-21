package flexiroute

import (
	"fmt"
	"net/http"
)

func BuildStdLibHandler(opts ...RouteSpecOpts) http.Handler {
	rs := defaultRouteSpec("")

	for _, opt := range opts {
		opt(rs)
	}

	rs.loadTemplates()

	mux := http.NewServeMux()
	buildRoutes(rs, mux)
	return mux
}

func buildRoutes(rs *RouteSpecification, mux *http.ServeMux) {
	for _, route := range rs.routes {
		for _, mw := range rs.middlewares {
			route.WithMiddleWare(mw)
		}

		path := fmt.Sprintf("%v%v", rs.name, route.Path())
		if rs.name != "" {
			path = fmt.Sprintf("/%v%v", rs.name, route.Path())
		}

		mux.HandleFunc(path, route.ToHandlerFunc(rs.templates))
	}

	for _, ns := range rs.subRoutes {
		buildRoutes(ns, mux)
	}
}
