package main

import (
	"net/http"

	"github.com/aganomaminmi/go-mvc/pkg/controller"
	"github.com/go-chi/chi/v5"
)

func NewRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome."))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm good."))
	})

	r.Route("/users", func(r chi.Router) {

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			controller.CreateUser(w, r)
		})

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				usrID := chi.URLParam(r, "userID")
				controller.GetUser(w, usrID)
			})
		})
	})
}
