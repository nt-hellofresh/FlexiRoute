package flexiroute

import (
	"github.com/stretchr/testify/assert"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFlexiRouteContext(t *testing.T) {
	t.Run("request details", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/users", nil)
		response := httptest.NewRecorder()

		ctx := NewContext(WithReaderWriter(response, request))

		assert.Equal(t, http.MethodPost, ctx.RequestMethod())
		assert.Equal(t, "/users", ctx.RequestURL().Path)
	})

	t.Run("render templates", func(t *testing.T) {

		tpl, err := template.New("foo").Parse(
			`{{define "home"}}<html>Hello, {{.Name}}!</html>{{end}}`,
		)

		assert.NoError(t, err)

		t.Run("template found", func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/users", nil)
			response := httptest.NewRecorder()

			ctx := NewContext(
				WithReaderWriter(response, request),
				WithHtmlTemplate(tpl),
			)

			type value struct {
				Name string
			}

			err = ctx.Render("home", value{Name: "World"})

			assert.NoError(t, err)
			assert.Equal(t, "<html>Hello, World!</html>", response.Body.String())
		})

		t.Run("template not found", func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/users", nil)
			response := httptest.NewRecorder()

			ctx := NewContext(
				WithReaderWriter(response, request),
				WithHtmlTemplate(tpl),
			)

			err = ctx.Render("unknown", nil)

			assert.EqualError(t, err, `html/template: "unknown" is undefined`)
		})

	})
}
