package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wxvn/go-clicker-ws/internal/domain"
	"github.com/wxvn/go-clicker-ws/internal/feature/auth/repository"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, "validation failed")
		return
	}

	auth := domain.Auth{
		Username: req.Username,
		Password: req.Password,
	}

	user, sessionID, err := h.service.Register(r.Context(), auth)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrUsernameConflict):
			writeError(w, http.StatusConflict, "username already exists")

		default:
			writeError(w, http.StatusInternalServerError, "internal server error")
		}

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   60 * 60 * 24 * 7,
		SameSite: http.SameSiteLaxMode,
	})

	res := AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Clicks:   user.Clicks,
	}

	WriteJSONResponse(w, http.StatusOK, res)
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})
}
