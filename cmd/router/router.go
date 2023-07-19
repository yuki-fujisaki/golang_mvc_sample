package router

import (
	"golang_mvc_sample/pkg/controller"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", controller.GetAllUsers)
		r.Post("/", controller.CreateUser)
	})

	return r
}
