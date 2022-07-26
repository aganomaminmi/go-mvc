package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
  r := chi.NewRouter()
  r.Use(middleware.RequestID)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  NewRoutes(r)

  if err := http.ListenAndServe(":9000", r); err !=nil {
    log.Fatal(err)
  }
}
