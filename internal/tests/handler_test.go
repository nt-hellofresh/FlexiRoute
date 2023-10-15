package tests

import (
	"encoding/json"
	"errors"
	"github.com/nt-hellofresh/flexiroute/internal/routes"
	"github.com/nt-hellofresh/flexiroute/pkg/flexiroute"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func testHandler(t *testing.T, handler flexiroute.ApiHandler) *routeWrapper {
	t.Helper()
	return &routeWrapper{
		route: flexiroute.NewApiRoute("/", handler, http.MethodGet),
	}
}

func TestFlexiRouteHandlers(t *testing.T) {
	t.Run("GetUsersHandler", func(t *testing.T) {
		handler := testHandler(t, routes.GetUsersHandler)
		resp := handler.get(t, "/")

		assert.Equal(t, http.StatusOK, resp.Code)

		type user struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		var result []user
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&result))

		expected := []user{
			{Name: "Teddy", Age: 24},
			{Name: "Sally", Age: 20},
			{Name: "Herschel", Age: 52},
			{Name: "Mary", Age: 39},
		}

		assert.Equal(t, expected, result)
	})
	t.Run("InternalServerError", func(t *testing.T) {
		handler := testHandler(t, func(ctx *flexiroute.Context) error {
			return errors.New("something unexpected went wrong")
		})

		resp := handler.get(t, "/")

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
