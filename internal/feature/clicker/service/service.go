package service

import (
	"context"
	"fmt"

	"github.com/wxvn/go-clicker-ws/internal/domain"
)

type ClickRepository interface {
	GetUserByID(ctx context.Context, id string) (domain.User, error)
	IncrementClicks(ctx context.Context, id string) (int64, error)
	GetLeaderboard(ctx context.Context, limit int, userID string) (domain.LeaderboardResult, error)
}

type ClickService struct {
	clickRepository ClickRepository
}

func NewClickService(clickRepository ClickRepository) *ClickService {
	return &ClickService{
		clickRepository: clickRepository,
	}
}

func (s *ClickService) IncrementClicks(ctx context.Context, userID string) (int64, error) {
	click, err := s.clickRepository.IncrementClicks(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("IncrementClicks: %w", err)
	}

	return click, nil
}

func (s *ClickService) GetUser(ctx context.Context, userId string) (domain.User, error) {
	user, err := s.clickRepository.GetUserByID(ctx, userId)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	return user, nil

}

func (s *ClickService) GetLeaderboard(ctx context.Context, limit int, userID string) (domain.LeaderboardResult, error) {
	leaderboard, err := s.clickRepository.GetLeaderboard(ctx, limit, userID)
	if err != nil {
		return domain.LeaderboardResult{}, fmt.Errorf("get leaderboard: %w", err)
	}

	return leaderboard, nil
}
