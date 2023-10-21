package flexiroute

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

type RouteSpecOpts func(rs *RouteSpecification)

type RouteSpecification struct {
	name         string
	routes       []*ApiRoute
	middlewares  []Middleware
	subRoutes    []*RouteSpecification
	templatesDir string
	templates    *template.Template
}

func (rs *RouteSpecification) loadTemplates() {
	templates, err := template.ParseGlob(rs.templatesDir)

	if err != nil {
		log.Fatalf("Failed to load templates, %v\n", err)
	}

	rs.templates = templates

	for _, ns := range rs.subRoutes {
		ns.templates = templates
	}
}

func defaultRouteSpec(name string) *RouteSpecification {
	return &RouteSpecification{
		name: name,
	}
}

func Use(mw Middleware) RouteSpecOpts {
	return func(rs *RouteSpecification) {
		rs.middlewares = append(rs.middlewares, mw)
	}
}

func Get(pattern string, handler ApiHandler) RouteSpecOpts {
	return withRoute(http.MethodGet, pattern, handler)
}

func Put(pattern string, handler ApiHandler) RouteSpecOpts {
	return withRoute(http.MethodPut, pattern, handler)
}

func Namespace(name string, opts ...RouteSpecOpts) RouteSpecOpts {
	return func(rs *RouteSpecification) {
		ns := defaultRouteSpec(name)

		for _, opt := range opts {
			opt(ns)
		}

		rs.subRoutes = append(rs.subRoutes, ns)
	}
}

func HtmlTemplatesDir(directory string) RouteSpecOpts {
	return func(rs *RouteSpecification) {
		if rs.name != "" {
			log.Fatal(errors.New("must only configure templates with the root handler"))
		}
		rs.templatesDir = directory
	}
}

func withRoute(method, pattern string, handler ApiHandler) RouteSpecOpts {
	return func(rs *RouteSpecification) {
		rs.routes = append(rs.routes, NewApiRoute(pattern, handler, method))
	}
}
