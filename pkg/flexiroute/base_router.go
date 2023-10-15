package flexiroute

import (
	"html/template"
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
	r.addRoute(pattern, handler, GET)
}

func (r *BaseRouter) Put(pattern string, handler ApiHandler) {
	r.addRoute(pattern, handler, PUT)
}

func (r *BaseRouter) Patch(pattern string, handler ApiHandler) {
	r.addRoute(pattern, handler, PATCH)
}

func (r *BaseRouter) Post(pattern string, handler ApiHandler) {
	r.addRoute(pattern, handler, POST)
}

func (r *BaseRouter) Delete(pattern string, handler ApiHandler) {
	r.addRoute(pattern, handler, DELETE)
}

func (r *BaseRouter) addRoute(pattern string, handler ApiHandler, method HttpMethod) {
	r.routes = append(r.routes, NewApiRoute(pattern, handler, method))
}
