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
