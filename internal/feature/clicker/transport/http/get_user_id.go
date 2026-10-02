package http

import "net/http"

func (h *ClicksHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")
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
