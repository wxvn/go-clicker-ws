package http

import (
	"net/http"

	"github.com/wxvn/go-clicker-ws/internal/transport/http/middleware"
)

func (h *ClicksHandler) GetMeUser(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	user, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid username or password")
		return
	}

	res := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Clicks:   user.Clicks,
	}

	WriteJSONResponse(w, http.StatusOK, res)

}
