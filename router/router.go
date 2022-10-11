package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shoot_and_win/service"
	"github.com/shoot_and_win/handler"
	"github.com/shoot_and_win/hub"
)

func SetupRouter(r chi.Router, s service.Service, h *hub.Hub) {
	r.Use(middleware.Logger)
	r.Post("/register_player", handler.RegisterPlayerHandler(s))
	r.Get("/ws", handler.WebsocketHandler(h))
}
