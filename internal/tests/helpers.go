package tests

import (
	"github.com/nt-hellofresh/flexiroute/pkg/flexiroute"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testServer struct {
	handler     http.Handler
	description string
}

func (s testServer) get(target string) *httptest.ResponseRecorder {
	return s.request(http.MethodGet, target)
}

func (s testServer) put(target string) *httptest.ResponseRecorder {
	return s.request(http.MethodPut, target)
}

func (s testServer) request(method, target string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(method, target, nil)

	s.handler.ServeHTTP(w, r)

	return w
}

type routeWrapper struct {
	route *flexiroute.ApiRoute
}

func (rw *routeWrapper) get(t *testing.T, target string) *httptest.ResponseRecorder {
	t.Helper()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, target, nil)

	rw.route.ToHandlerFunc(nil).ServeHTTP(w, r)

	return w
}
