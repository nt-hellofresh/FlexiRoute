package flexiroute

import (
	"html/template"
	"net/http"
)

type BaseRouter struct {
	name        string
	routes      []*ApiRoute
	middlewares []Middleware
	templates   *template.Template
}

func (r *BaseRouter) Use(fn Middleware) {
	r.middlewares = append(r.middlewares, fn)
}

func (r *BaseRouter) Get(pattern string, handler ApiHandler) {
	r.addRoute(http.MethodGet, pattern, handler)
}

func (r *BaseRouter) Put(pattern string, handler ApiHandler) {
	r.addRoute(http.MethodPut, pattern, handler)
}

func (r *BaseRouter) Patch(pattern string, handler ApiHandler) {
	r.addRoute(http.MethodPatch, pattern, handler)
}

func (r *BaseRouter) Post(pattern string, handler ApiHandler) {
	r.addRoute(http.MethodPost, pattern, handler)
}

func (r *BaseRouter) Delete(pattern string, handler ApiHandler) {
	r.addRoute(http.MethodDelete, pattern, handler)
}

func (r *BaseRouter) addRoute(method, pattern string, handler ApiHandler) {
	r.routes = append(r.routes, NewApiRoute(pattern, handler, method))
}
