package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	repo := NewInMemoryRepository()
	s := NewService(repo)
	
	r := chi.NewRouter()
	
	setupRouter(r, s)

	http.ListenAndServe("localhost:8080", r)

}