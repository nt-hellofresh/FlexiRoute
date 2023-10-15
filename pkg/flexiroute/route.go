package flexiroute

import (
	"html/template"
	"net/http"
)

type ApiRoute struct {
	pattern string
	handler ApiHandler
	method  HttpMethod
}

func NewApiRoute(pattern string, handler ApiHandler, method HttpMethod) *ApiRoute {
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

func (route *ApiRoute) Method() HttpMethod {
	return route.method
}

func (route *ApiRoute) ToHandlerFunc(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(
			WithReaderWriter(w, r),
			WithHtmlTemplate(templates),
		)

		if ok := ctx.AllowMethod(route.method); !ok {
			return
		}

		if err := route.handler(ctx); err != nil {
			errMsg := map[string]string{"error": "An error occurred", "detail": err.Error()}
			_ = ctx.WriteJSON(http.StatusInternalServerError, errMsg)
		}
	}
}
