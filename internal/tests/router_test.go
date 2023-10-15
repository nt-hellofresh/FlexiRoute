package tests

import (
	"github.com/nt-hellofresh/flexiroute/internal"
	"github.com/nt-hellofresh/flexiroute/pkg/flexiroute"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func testCase(t *testing.T, fn func() flexiroute.RouterFacade, description string) testServer {
	t.Helper()
	router := fn()
	internal.Configure(router)
	return testServer{
		handler:     router,
		description: description,
	}
}

func TestFlexiRouter(t *testing.T) {
	for _, server := range []testServer{
		testCase(t, flexiroute.NewDefaultRouter, "using standard library router facade"),
		testCase(t, flexiroute.NewChiRouter, "using chi router facade"),
	} {
		t.Run(server.description, func(t *testing.T) {
			t.Run("requesting index", func(t *testing.T) {
				resp := server.get("/")
				assert.Equal(t, http.StatusOK, resp.Code)
			})
			t.Run("requesting dice", func(t *testing.T) {
				resp := server.get("/dice")
				assert.Equal(t, http.StatusOK, resp.Code)
			})
			t.Run("requesting users", func(t *testing.T) {
				resp := server.get("/users/")
				assert.Equal(t, http.StatusOK, resp.Code)

				t.Run("test GET", func(t *testing.T) {
					getResp := server.get("/users/test")
					assert.Equal(t, http.StatusMethodNotAllowed, getResp.Code)
				})
				t.Run("test PUT", func(t *testing.T) {
					putResp := server.put("/users/test")
					assert.Equal(t, http.StatusOK, putResp.Code)
				})
			})
		})
	}
}
