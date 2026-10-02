package http

import (
	"net/http"

	"github.com/wxvn/go-clicker-ws/internal/transport/http/middleware"
)

func (h *ClicksHandler) Click(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	clicks, err := h.service.IncrementClicks(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusNotModified, "invalid username or password")
		return
	}

	res := ClickResponse{
		Clicks: clicks,
	}

	WriteJSONResponse(w, http.StatusOK, res)

}
