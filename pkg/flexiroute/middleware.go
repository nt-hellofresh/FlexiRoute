package flexiroute

import "net/http"

type Middleware func(handler ApiHandler) ApiHandler

func (mw Middleware) ToHandler() func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		apiHandler := mw(func(ctx *Context) error {
			handler.ServeHTTP(ctx.w, ctx.r)
			return nil
		})

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			f := apiHandler.ToHandlerFunc(r.Method, nil)
			f(w, r)
		})
	}
}
