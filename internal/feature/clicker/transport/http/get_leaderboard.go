package http

import (
	"net/http"

	"github.com/wxvn/go-clicker-ws/internal/transport/http/middleware"
)

func (h *ClicksHandler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	leaderboard, err := h.service.GetLeaderboard(r.Context(), 10, userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to get leaderboard")
		return
	}

	response := LeaderboardResponse{
		Users: make([]LeaderboardUserResponse, 0, len(leaderboard.Users)),
	}

	for _, user := range leaderboard.Users {
		response.Users = append(response.Users, LeaderboardUserResponse{
			ID:       user.ID,
			Username: user.Username,
			Avatar:   user.Avatar,
			Clicks:   user.Clicks,
			Position: user.Position,
		})
	}

	if leaderboard.CurrentUser != nil {
		user := leaderboard.CurrentUser

		response.CurrentUser = &LeaderboardUserResponse{
			ID:       user.ID,
			Username: user.Username,
			Avatar:   user.Avatar,
			Clicks:   user.Clicks,
			Position: user.Position,
		}
	}

	WriteJSONResponse(w, http.StatusOK, response)
}
