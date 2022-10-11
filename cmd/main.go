package main

import (
	"net/http"

	"github.com/shoot_and_win/repository"
	"github.com/shoot_and_win/service"
	"github.com/shoot_and_win/hub"
	"github.com/shoot_and_win/router"

	"github.com/go-chi/chi/v5"
)

func main() {
	repo := repository.NewInMemoryRepository()
	s := service.NewService(repo)
	h := hub.NewHub()
	r := chi.NewRouter()

	router.SetupRouter(r, s, h)

	http.ListenAndServe("localhost:8080", r)
}
