package flexiroute

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFlexiRouteHandler(t *testing.T) {
	var apiHandler ApiHandler = func(ctx *Context) error {
		if ctx.RequestURL().String() == "/throw-error" {
			return errors.New("simulated error")
		}
		return ctx.WriteJSON(http.StatusOK, nil)
	}

	handler := apiHandler.ToHandlerFunc(http.MethodGet, nil)

	t.Run("handles request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		handler(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("method not allowed", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()

		handler(response, request)

		assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/throw-error", nil)
		response := httptest.NewRecorder()

		handler(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}
