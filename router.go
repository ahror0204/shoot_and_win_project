package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setupRouter(r chi.Router, s Service) {
	r.Use(middleware.Logger)
	r.Post("/register_player", RegisterPlayerHandler(s))
}
