package main

import (
	"fmt"
	"net/http"
	"simpleCrudGo/configs"
	"simpleCrudGo/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
	r.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetAPIPort()), r)
}
