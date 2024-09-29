package routes

import (
	"github.com/aguradev/react-go-auth/internal/delivery/http/handler"
	"github.com/go-chi/chi/v5"
)

type AuthRouteConfig struct {
	Router      *chi.Mux
	AuthHandler handler.AuthHandlerInterface
}

func (h *AuthRouteConfig) AuthRoute() {
	h.Router.Get("/", h.AuthHandler.Authenticable)
}
