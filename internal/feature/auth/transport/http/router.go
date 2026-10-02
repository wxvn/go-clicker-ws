package http

import (
	"net/http"
)

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /auth/register", h.Register)
	router.HandleFunc("POST /auth/login", h.Login)
	router.HandleFunc("POST /auth/logout", h.Logout)

}
