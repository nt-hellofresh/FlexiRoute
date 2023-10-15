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

func NewDefaultRouter() *DefaultRouterFacade {
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

func (r *DefaultRouterFacade) buildRoutes() {
	for _, route := range r.routes {
		for _, mw := range r.middlewares {
			route.WithMiddleWare(mw)
		}

		path := fmt.Sprintf("%v%v", r.name, route.Path())
		if r.name != "" {
			path = fmt.Sprintf("/%v%v", r.name, route.Path())
		}

		http.HandleFunc(path, route.ToHandlerFunc(r.templates))
	}

	for _, ns := range r.namespaces {
		ns.buildRoutes()
	}
}

func (r *DefaultRouterFacade) ServeHTTP(addr string) error {
	r.buildRoutes()
	return http.ListenAndServe(addr, nil)
}
