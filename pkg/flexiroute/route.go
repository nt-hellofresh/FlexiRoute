package flexiroute

import (
	"html/template"
	"net/http"
)

type ApiRoute struct {
	pattern string
	handler ApiHandler
	method  string
}

func NewApiRoute(pattern string, handler ApiHandler, method string) *ApiRoute {
	return &ApiRoute{
		pattern: pattern,
		handler: handler,
		method:  method,
	}
}

func (route *ApiRoute) WithMiddleWare(mw Middleware) {
	route.handler = mw(route.handler)
}

func (route *ApiRoute) Path() string {
	return route.pattern
}

func (route *ApiRoute) Method() string {
	return route.method
}

func (route *ApiRoute) ToHandlerFunc(templates *template.Template) http.HandlerFunc {
	return route.handler.ToHandlerFunc(route.method, templates)
}
