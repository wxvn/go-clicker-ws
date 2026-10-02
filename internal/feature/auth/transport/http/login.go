package http

import (
	"encoding/json"
	"net/http"

	"github.com/wxvn/go-clicker-ws/internal/domain"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := DecodeAndValidateRequest(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request")
		return
	}
	auth := domain.Auth{
		Username: req.Username,
		Password: req.Password,
	}

	user, sessionID, err := h.service.Login(r.Context(), auth)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid username or password")
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

func WriteJSONResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
