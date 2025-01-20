package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) run() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Ok\n"))
		})

		r.Route("/users", func(r chi.Router) {
			// r.Get("/", app.findAll)
			r.Post("/create", app.createUser)
			r.Get("/{id}", app.findById)
			r.Post("/login", app.login)
		})

	})

	return r
}
