package flexiroute

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

type ChiRouterFacade struct {
	BaseRouter
	namespaces []*ChiRouterFacade
}

func NewChiRouter() *ChiRouterFacade {
	return &ChiRouterFacade{}
}

func (r *ChiRouterFacade) Namespace(name string) RouterFacade {
	ns := &ChiRouterFacade{
		BaseRouter: BaseRouter{
			name:      name,
			templates: r.templates,
		},
	}
	r.namespaces = append(r.namespaces, ns)
	return ns
}

func (r *ChiRouterFacade) LoadTemplates(directory string) {
	templates, err := template.ParseGlob(directory)

	if err != nil {
		log.Fatalf("Failed to load templates, %v\n", err)
	}

	r.templates = templates

	for _, ns := range r.namespaces {
		ns.templates = templates
	}
}

func (r *ChiRouterFacade) ServeHTTP(addr string) error {
	app := chi.NewRouter()
	r.buildRoutes(app)
	return http.ListenAndServe(addr, app)
}

func (r *ChiRouterFacade) buildRoutes(app *chi.Mux) {
	for _, route := range r.routes {
		for _, mw := range r.middlewares {
			route.WithMiddleWare(mw)
		}

		path := route.Path()

		if r.name != "" {
			// Slightly cheating here as proper nested chi routes aren't
			// really utilised. Instead, the fully qualified path is
			// registered with the application.
			path = fmt.Sprintf("/%v%v", r.name, route.Path())
		}

		r.registerEndpoint(app, path, route)
	}

	for _, ns := range r.namespaces {
		ns.buildRoutes(app)
	}
}

func (r *ChiRouterFacade) registerEndpoint(app chi.Router, path string, route *ApiRoute) {
	handler := route.ToHandlerFunc(r.templates)
	switch route.Method() {
	case GET:
		app.Get(path, handler)
	case PUT:
		app.Put(path, handler)
	case PATCH:
		app.Patch(path, handler)
	case POST:
		app.Post(path, handler)
	case DELETE:
		app.Delete(path, handler)
	default:
		log.Fatalf("unsupported method")
	}
}
