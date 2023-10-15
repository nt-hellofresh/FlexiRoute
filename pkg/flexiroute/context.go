package flexiroute

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type ContextOpts func(ctx *Context)

type Context struct {
	w         http.ResponseWriter
	r         *http.Request
	templates *template.Template
}

func WithReaderWriter(w http.ResponseWriter, r *http.Request) ContextOpts {
	return func(ctx *Context) {
		ctx.w = w
		ctx.r = r
	}
}

func WithHtmlTemplate(templates *template.Template) ContextOpts {
	return func(ctx *Context) {
		ctx.templates = templates
	}
}

func NewContext(opts ...ContextOpts) *Context {
	ctx := &Context{}

	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}

func (ctx *Context) Render(templateName string, value any) error {
	return ctx.templates.ExecuteTemplate(ctx.w, templateName, value)
}

func (ctx *Context) RequestMethod() string {
	return ctx.r.Method
}

func (ctx *Context) RequestURL() *url.URL {
	return ctx.r.URL
}

func (ctx *Context) WriteJSON(statusCode int, v any) error {
	ctx.w.Header().Add("Content-Type", "application/json")
	ctx.w.WriteHeader(statusCode)
	return json.NewEncoder(ctx.w).Encode(v)
}

func (ctx *Context) AllowMethod(method HttpMethod) bool {
	if ctx.r.Method != string(method) {
		errMsg := map[string]string{
			"error":  fmt.Sprintf("%v method not allowed", ctx.r.Method),
			"detail": fmt.Sprintf("Only %v method is allowed", method),
		}
		_ = ctx.WriteJSON(http.StatusBadRequest, errMsg)
		return false
	}
	return true
}
