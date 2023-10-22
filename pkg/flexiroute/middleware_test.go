package flexiroute

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	t.Run("wraps handler", func(t *testing.T) {
		isMiddlewareCalled := false
		var mw Middleware = func(h ApiHandler) ApiHandler {
			isMiddlewareCalled = true
			return h
		}

		isHandlerCalled := false
		var handler ApiHandler = func(ctx *Context) error {
			isHandlerCalled = true
			return ctx.WriteJSON(http.StatusOK, nil)
		}

		handler = mw(handler)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		assert.NoError(t, handler(NewContext(WithReaderWriter(response, request))))
		assert.True(t, isHandlerCalled, "handler should be called")
		assert.True(t, isMiddlewareCalled, "middleware should be called")
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("converts to http.handler", func(t *testing.T) {
		isMiddlewareCalled := false
		var mw Middleware = func(h ApiHandler) ApiHandler {
			isMiddlewareCalled = true
			return h
		}

		isHandlerCalled := false
		var handler ApiHandler = func(ctx *Context) error {
			isHandlerCalled = true
			return ctx.WriteJSON(http.StatusOK, nil)
		}

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		f := handler.ToHandlerFunc(http.MethodGet, nil)
		h := http.Handler(f)
		h = mw.ToHandler()(h)

		h.ServeHTTP(response, request)

		assert.True(t, isHandlerCalled, "handler should be called")
		assert.True(t, isMiddlewareCalled, "middleware should be called")
		assert.Equal(t, http.StatusOK, response.Code)
	})
}
