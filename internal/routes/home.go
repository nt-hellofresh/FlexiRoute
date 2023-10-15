package routes

import (
	"github.com/nt-hellofresh/flexiroute/pkg/flexiroute"
	"math/rand"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUsers() []user {
	return []user{
		{Name: "Teddy", Age: 24},
		{Name: "Sally", Age: 20},
		{Name: "Herschel", Age: 52},
		{Name: "Mary", Age: 39},
	}
}

func GetUsersHandler(ctx *flexiroute.Context) error {
	resp := getUsers()
	return ctx.WriteJSON(http.StatusOK, resp)
}

func RandomNumberHandler(ctx *flexiroute.Context) error {
	type data struct {
		Value int `json:"value"`
	}
	return ctx.WriteJSON(200, data{rand.Intn(6) + 1})
}

func HomeHandler(ctx *flexiroute.Context) error {
	type data struct {
		Name string
	}
	return ctx.Render("home.html", data{Name: "World"})
}

func PutHandler(ctx *flexiroute.Context) error {
	resp := map[string]string{"response": "This is a PUT only endpoint!"}
	return ctx.WriteJSON(http.StatusOK, resp)
}
