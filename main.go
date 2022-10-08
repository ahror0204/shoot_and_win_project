package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	repo := NewInMemoryRepository()
	s := NewService(repo)
	h := NewHub()
	r := chi.NewRouter()
	
	setupRouter(r, s, h)

	http.ListenAndServe("localhost:8080", r)

}