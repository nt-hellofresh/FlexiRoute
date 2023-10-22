package flexiroute

import (
	"html/template"
	"net/http"
)

type ApiHandler func(ctx *Context) error

func (h ApiHandler) ToHandlerFunc(method string, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(
			WithReaderWriter(w, r),
			WithHtmlTemplate(templates),
		)

		if ok := ctx.AllowMethod(method); !ok {
			return
		}

		if err := h(ctx); err != nil {
			errMsg := map[string]string{"error": "An error occurred", "detail": err.Error()}
			_ = ctx.WriteJSON(http.StatusInternalServerError, errMsg)
		}
	}
}

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
