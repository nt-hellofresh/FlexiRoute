package flexiroute

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type DefaultRouterFacade struct {
	BaseRouter
	namespaces []*DefaultRouterFacade
}

func NewDefaultRouter() RouterFacade {
	return &DefaultRouterFacade{}
}

func (r *DefaultRouterFacade) Namespace(name string) RouterFacade {
	ns := &DefaultRouterFacade{
		BaseRouter: BaseRouter{
			name:      name,
			templates: r.templates,
		},
	}
	r.namespaces = append(r.namespaces, ns)
	return ns
}

func (r *DefaultRouterFacade) LoadTemplates(directory string) {
	templates, err := template.ParseGlob(directory)

	if err != nil {
		log.Fatalf("Failed to load templates, %v\n", err)
	}

	r.templates = templates

	for _, ns := range r.namespaces {
		ns.templates = templates
	}
}

func (r *DefaultRouterFacade) buildRoutes(mux *http.ServeMux) {
	for _, route := range r.routes {
		for _, mw := range r.middlewares {
			route.WithMiddleWare(mw)
		}

		path := fmt.Sprintf("%v%v", r.name, route.Path())
		if r.name != "" {
			path = fmt.Sprintf("/%v%v", r.name, route.Path())
		}

		mux.HandleFunc(path, route.ToHandlerFunc(r.templates))
	}

	for _, ns := range r.namespaces {
		ns.buildRoutes(mux)
	}
}

func (r *DefaultRouterFacade) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mux := http.NewServeMux()
	r.buildRoutes(mux)
	mux.ServeHTTP(w, req)
}
