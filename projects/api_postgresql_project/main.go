package main

import (
	"fmt"
	"net/http"

	"github.com/GuiCintra27/go/projects/api_postgresql_project/configs"
	"github.com/GuiCintra27/go/projects/api_postgresql_project/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)

	fmt.Println("Server running on port", configs.GetAPI().Port)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetAPI().Port), r)
}